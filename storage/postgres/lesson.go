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

type lessonRepo struct {
	db *pgxpool.Pool
}

func NewLesson(db *pgxpool.Pool) lessonRepo {
	return lessonRepo{
		db: db,
	}
}

func (c *lessonRepo) Create(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {

	id := uuid.New()
	query := `INSERT INTO "lesson" (
		id,
		schedule_id,
		group_id,
		from,
		to,
		theme,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		lesson.ScheduleId,
		lesson.GroupId,
		lesson.From,
		lesson.To,
		lesson.Theme,
	)

	if err != nil {
		return models.Lesson{}, err
	}
	return models.Lesson{
		Id:         lesson.Id,
		ScheduleId: lesson.ScheduleId,
		GroupId:    lesson.GroupId,
		From:       lesson.From,
		To:         lesson.To,
		Theme:      lesson.Theme,
		Created_at: lesson.Created_at,
		Updated_at: lesson.Updated_at,
	}, nil
}

func (c *lessonRepo) Update(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {
	query := `update "lesson" set 
	schedule_id=$1,
	group_id=$2,
	from=$3,
	to=$4,
	theme=$5,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $6
	`
	_, err := c.db.Exec(context.Background(), query,
		lesson.ScheduleId,
		lesson.GroupId,
		lesson.From,
		lesson.To,
		lesson.Theme,
		lesson.Id,
	)
	if err != nil {
		return models.Lesson{}, err
	}
	return models.Lesson{
		Id:         lesson.Id,
		ScheduleId: lesson.ScheduleId,
		GroupId:    lesson.GroupId,
		From:       lesson.From,
		To:         lesson.To,
		Theme:      lesson.Theme,
		Created_at: lesson.Created_at,
		Updated_at: lesson.Updated_at,
	}, nil
}

func (c *lessonRepo) GetAll(ctx context.Context, req models.GetAllLessonsRequest) (models.GetAllLessonsResponse, error) {
	var (
		resp   = models.GetAllLessonsResponse{}
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
        schedule_id,
        group_id,
		from,
		to,
		theme,
        created_at,
        updated_at
        FROM "lesson"`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			lesson      = models.Lesson{}
			schedule_id sql.NullString
			group_id    sql.NullString
			from        sql.NullString
			to          sql.NullString
			theme       sql.NullString
			created_at  sql.NullString
			updateAt    sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&lesson.Id,
			&schedule_id,
			&group_id,
			&from,
			&to,
			&theme,
			&created_at,
			&updateAt); err != nil {
			return resp, err
		}
		lesson.Updated_at = pkg.NullStringToString(updateAt)
		resp.Lessons = append(resp.Lessons, models.Lesson{
			ScheduleId: schedule_id.String,
			GroupId:    group_id.String,
			From:       from.String,
			To:         to.String,
			Created_at: created_at.String,
			Updated_at: updateAt.String,
		})
	}
	return resp, nil
}

func (c *lessonRepo) GetByID(ctx context.Context, id string) (models.Lesson, error) {
	var (
		lesson      = models.Lesson{}
		schedule_id sql.NullString
		group_id    sql.NullString
		from        sql.NullString
		to          sql.NullString
		theme       sql.NullString
		created_at  sql.NullString
		updateAt    sql.NullString
	)

	if err := c.db.QueryRow(context.Background(), `select id, schedule_id, group_id, from, to, theme, created_at, updated_at from "lesson" where id = $1`, id).Scan(
		&lesson.Id,
		&schedule_id,
		&group_id,
		&from,
		&to,
		&theme,
		&created_at,
		&updateAt); err != nil {
		return models.Lesson{}, err
	}
	return models.Lesson{
		ScheduleId: schedule_id.String,
		GroupId:    group_id.String,
		From:       from.String,
		To:         to.String,
		Created_at: created_at.String,
		Updated_at: updateAt.String,
	}, nil
}

func (c *lessonRepo) Delete(ctx context.Context, id string) error {
	query := `delete from "lesson" where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
