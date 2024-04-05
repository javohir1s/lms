package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"lms_back/api/models"
	"lms_back/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepo struct {
	db *pgxpool.Pool
}

func NewTask(db *pgxpool.Pool) TaskRepo {
	return TaskRepo{
		db: db,
	}
}

func (c *TaskRepo) Create(ctx context.Context,task models.Task) (models.Task, error) {

	id := uuid.New()
	query := `INSERT INTO "task" (
		id,
		lesson_id
		group_id,
		score,
		created_at)
		VALUES($1,$2,$3,$4,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		task.LessonId,
		task.GroupId,
		task.Score,
	)

	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		Id:        task.Id,
		LessonId:  task.LessonId,
		GroupId:   task.GroupId,
		Score:     task.Score,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (c *TaskRepo) Update(ctx context.Context,task models.Task) (models.Task, error) {
	query := `update "task" set
	lesson_id=$1, 
	group_id=$2,
    score=$3,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $4
	`
	_, err := c.db.Exec(context.Background(), query,
		task.LessonId,
		task.GroupId,
		task.Score,
		task.Id,
	)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		Id:        task.Id,
		LessonId:  task.LessonId,
		GroupId:   task.GroupId,
		Score:     task.Score,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (c *TaskRepo) GetAll(ctx context.Context,req models.GetAllTasksRequest) (models.GetAllTasksResponse, error) {
	var (
		resp   = models.GetAllTasksResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := c.db.Query(context.Background(), `select count(id) over(),
        id,
		lesson_id,
        group_id,
		score,
        created_at,
        updated_at
        FROM "task"`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			task       = models.Task{}
			lesson_id  sql.NullString
			group_id   sql.NullString
			created_at sql.NullString
			updated_at sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&task.Id,
			&lesson_id,
			&group_id,
			&created_at,
			&updated_at); err != nil {
			return resp, err
		}
		task.UpdatedAt = pkg.NullStringToString(updated_at)
		resp.Tasks = append(resp.Tasks, models.Task{
			LessonId:  lesson_id.String,
			GroupId:   group_id.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}
	return resp, nil
}

func (c *TaskRepo) GetByID(ctx context.Context, id string) (models.Task, error) {
	var (
		task       = models.Task{}
		lesson_id  sql.NullString
		group_id   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	if err := c.db.QueryRow(context.Background(), `select id, lesson_id, group_id, score, created_at, updated_at from "task" where id = $1`, id).Scan(
		&task.Id,
		&lesson_id,
		&group_id,
		&created_at,
		&updated_at); err != nil {
		return models.Task{}, err
	}
	return models.Task{
		LessonId:  lesson_id.String,
		GroupId:   group_id.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (c *TaskRepo) Delete(ctx context.Context, id string) error {
	query := `delete from "task" where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
