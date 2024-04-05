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

type ScheduleRepo struct {
	db *pgxpool.Pool
}

func NewSchedule(db *pgxpool.Pool) ScheduleRepo {
	return ScheduleRepo{
		db: db,
	}
}

func (c *ScheduleRepo) Create(ctx context.Context, schedule models.Schedule) (models.Schedule, error) {

	id := uuid.New()
	query := `INSERT INTO schedule (
		id,
		group_id,
		group_type,
		start_time,
		end_time,
		date,
		branch_id,
		teacher_id,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		schedule.Group_id,
		schedule.Group_type,
		schedule.Start_time,
		schedule.End_time,
		schedule.Date,
		schedule.Branch_id,
		schedule.Teacher_id,
		schedule.Created_at,
	)

	if err != nil {
		return models.Schedule{}, err
	}
	return models.Schedule{
		Id:         schedule.Id,
		Group_id:   schedule.Group_id,
		Group_type: schedule.Group_type,
		Start_time: schedule.Start_time,
		End_time:   schedule.End_time,
		Date:       schedule.Date,
		Branch_id:  schedule.Branch_id,
		Teacher_id: schedule.Teacher_id,
		Created_at: schedule.Created_at,
		Updated_at: schedule.Updated_at,
	}, nil
}

func (c *ScheduleRepo) Update(ctx context.Context, schedule models.Schedule) (models.Schedule, error) {
	query := `update schedule set 
	group_id=$1,
	group_type=$2,
	start_time=$3,
	end_time=$4,
	date=$5,
	branch_id=$6,
	teacher_id=$7,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $8
	`
	_, err := c.db.Exec(context.Background(), query,
		schedule.Group_id,
		schedule.Group_type,
		schedule.Start_time,
		schedule.End_time,
		schedule.Date,
		schedule.Branch_id,
		schedule.Teacher_id,
		schedule.Created_at,
		schedule.Id,
	)
	if err != nil {
		return models.Schedule{}, err
	}
	return models.Schedule{
		Id:         schedule.Id,
		Group_id:   schedule.Group_id,
		Group_type: schedule.Group_type,
		Start_time: schedule.Start_time,
		End_time:   schedule.End_time,
		Date:       schedule.Date,
		Branch_id:  schedule.Branch_id,
		Teacher_id: schedule.Teacher_id,
		Created_at: schedule.Created_at,
		Updated_at: schedule.Updated_at,
	}, nil
}

func (c *ScheduleRepo) GetAll(ctx context.Context, req models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error) {
	var (
		resp   = models.GetAllSchedulesResponse{}
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
		group_id,
		group_type,
		start_time,
		end_time,
		date,
		branch_id,
		teacher_id,
		created_at,
        updated_at FROM schedule`+filter+``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			schedule   = models.Schedule{}
			group_id   sql.NullString
			group_type sql.NullString
			start_time sql.NullString
			end_time   sql.NullString
			date       sql.NullString
			branch_id  sql.NullString
			teacher_id sql.NullString
			created_at sql.NullString
			updateAt   sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&schedule.Id,
			&group_id,
			&group_type,
			&start_time,
			&end_time,
			&date,
			&branch_id,
			&teacher_id,
			&created_at,
			&updateAt); err != nil {
			return resp, err
		}
		schedule.Updated_at = pkg.NullStringToString(updateAt)
		resp.Schedules = append(resp.Schedules, models.Schedule{
			Group_id:   group_id.String,
			Group_type: group_id.String,
			Start_time: start_time.String,
			End_time:   end_time.String,
			Date:       date.String,
			Branch_id:  branch_id.String,
			Teacher_id: teacher_id.String,
			Created_at: created_at.String,
			Updated_at: updateAt.String,
		})
	}
	return resp, nil
}

func (c *ScheduleRepo) GetByID(ctx context.Context, id string) (models.Schedule, error) {
	var (
		schedule   = models.Schedule{}
		group_id   sql.NullString
		group_type sql.NullString
		start_time sql.NullString
		end_time   sql.NullString
		date       sql.NullString
		branch_id  sql.NullString
		teacher_id sql.NullString
		created_at sql.NullString
		updateAt   sql.NullString
	)
	if err := c.db.QueryRow(context.Background(), `select id, group_id, group_type, start_time,
			end_time, date, branch_id, teacher_id, created_at,
			updated_at from schedule where id = $1`, id).Scan(
		&schedule.Id,
		&group_id,
		&group_type,
		&start_time,
		&end_time,
		&date,
		&branch_id,
		&teacher_id,
		&created_at,
		&updateAt,
	); err != nil {
		return models.Schedule{}, err
	}
	return models.Schedule{
		Group_id:   group_id.String,
		Group_type: group_id.String,
		Start_time: start_time.String,
		End_time:   end_time.String,
		Date:       date.String,
		Branch_id:  branch_id.String,
		Teacher_id: teacher_id.String,
		Created_at: created_at.String,
		Updated_at: updateAt.String,
	}, nil
}

func (c *ScheduleRepo) Delete(ctx context.Context, id string) error {
	query := `delete from schedule where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
