package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"lms_back/api/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacher(db *pgxpool.Pool) TeacherRepo {
	return TeacherRepo{
		db: db,
	}
}

func (c *TeacherRepo) Create(ctx context.Context, teacher models.Teacher) (models.Teacher, error) {

	id := uuid.New()
	query := `INSERT INTO teacher (
		id,
		full_name,
		email,
		age,
		status,
		login,
		password,
		created_at)
		VALUES($1,$2,$3,$4,$5,$6,$7,CURRENT_TIMESTAMP) 
	`
	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		teacher.Full_name,
		teacher.Email,
		teacher.Age,
		teacher.Status,
		teacher.Login,
		teacher.Password,
	)

	if err != nil {
		return models.Teacher{}, err
	}

	return models.Teacher{
		Id:         id.String(),
		Email:      teacher.Email,
		Age:        teacher.Age,
		Status:     teacher.Status,
		Login:      teacher.Login,
		Password:   teacher.Password,
		Created_at: teacher.Created_at,
		Updated_at: teacher.Updated_at,
	}, nil
}

func (c *TeacherRepo) Update(ctx context.Context, teacher models.Teacher) (models.Teacher, error) {
	query := `update teacher set 
	full_name=$1,
	email=$2,
	age=$3,
	login=$4,
	password=$5,
	status=$6,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $7 AND deleted_at = 0
	`
	_, err := c.db.Exec(context.Background(), query,
		teacher.Full_name,
		teacher.Email,
		teacher.Age,
		teacher.Login,
		teacher.Password,
		teacher.Status,
		teacher.Id,
	)
	if err != nil {
		return models.Teacher{}, err
	}
	return models.Teacher{
		Id:         teacher.Id,
		Email:      teacher.Email,
		Age:        teacher.Age,
		Status:     teacher.Status,
		Login:      teacher.Login,
		Password:   teacher.Password,
		Created_at: teacher.Created_at,
		Updated_at: teacher.Updated_at,
	}, nil
}

func (c *TeacherRepo) GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
	var (
		resp   = models.GetAllTeachersResponse{}
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
        full_name,
		email,
		age,
        login,
		password,
		status,
        created_at,
        updated_at,
        deleted_at FROM teacher WHERE deleted_at = 0`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			teacher    = models.Teacher{}
			full_name  sql.NullString
			email      sql.NullString
			age        sql.NullInt16
			login      sql.NullString
			password   sql.NullString
			status     sql.NullString
			created_at sql.NullString
			updated_at sql.NullString
			deleted_at sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&teacher.Id,
			&full_name,
			&email,
			&age,
			&login,
			&password,
			&status,
			&created_at,
			&updated_at,
			&deleted_at); err != nil {
			return resp, err
		}
		resp.Teachers = append(resp.Teachers, models.Teacher{
			Id:         teacher.Id,
			Full_name:  full_name.String,
			Email:      email.String,
			Age:        int(age.Int16),
			Login:      login.String,
			Password:   password.String,
			Status:     status.String,
			Created_at: created_at.String,
			Updated_at: updated_at.String,
			Deleted_at: deleted_at.String,
		})
	}
	return resp, nil
}

func (c *TeacherRepo) GetByID(ctx context.Context, id string) (models.Teacher, error) {
	teacher := models.Teacher{}
	var (
		full_name  sql.NullString
		email      sql.NullString
		age        sql.NullInt16
		login      sql.NullString
		password   sql.NullString
		status     sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
		deleted_at sql.NullString
	)

	if err := c.db.QueryRow(context.Background(), `select id, full_name, email, age, login, password, status, created_at, updated_at, deleted_at from teacher where id = $1`, id).Scan(
		&teacher.Id,
		&full_name,
		&email,
		&age,
		&login,
		&password,
		&status,
		&created_at,
		&updated_at,
		&deleted_at,
	); err != nil {
		return models.Teacher{}, err
	}
	return models.Teacher{
		Id:         teacher.Id,
		Full_name:  full_name.String,
		Email:      email.String,
		Age:        int(age.Int16),
		Login:      login.String,
		Password:   password.String,
		Status:     status.String,
		Created_at: created_at.String,
		Updated_at: updated_at.String,
		Deleted_at: deleted_at.String,
	}, nil
}

func (c *TeacherRepo) Delete(ctx context.Context, id string) error {
	query := `delete from teacher where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
