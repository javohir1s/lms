package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"lms_back/api/models"
	"lms_back/pkg"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
)

type adminRepo struct {
	db *pgxpool.Pool
}

func NewAdmin(db *pgxpool.Pool) adminRepo {
	return adminRepo{
		db: db,
	}
}

func (c *adminRepo) Create(ctx context.Context, admin models.Admin) (models.Admin, error) {

	id := uuid.New()
	query := `INSERT INTO "admin" (
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
		admin.Full_Name,
		admin.Email,
		admin.Age,
		admin.Status,
		admin.Login,
		admin.Password)

	if err != nil {
		return models.Admin{}, err
	}
	return models.Admin{
		Id:         id.String(),
		Full_Name:  admin.Full_Name,
		Email:      admin.Email,
		Age:        admin.Age,
		Status:     admin.Status,
		Login:      admin.Login,
		Password:   admin.Password,
		Created_at: admin.Created_at,
		Updated_at: admin.Updated_at,
	}, nil
}

func (c *adminRepo) Update(ctx context.Context, admin models.Admin) (models.Admin, error) {
	query := `update "admin" set 
	full_name=$1,
	email=$2,
	age=$3,
	status=$4,
	login=$5,
	password=$6,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $7
	`
	_, err := c.db.Exec(context.Background(), query,
		admin.Full_Name,
		admin.Email,
		admin.Age,
		admin.Status,
		admin.Login,
		admin.Password,
		admin.Id,
	)
	if err != nil {
		return models.Admin{}, err
	}
	return models.Admin{
		Id:         admin.Id,
		Full_Name:  admin.Email,
		Age:        admin.Age,
		Status:     admin.Status,
		Login:      admin.Login,
		Password:   admin.Password,
		Created_at: admin.Created_at,
		Updated_at: admin.Updated_at,
	}, nil
}

func (c *adminRepo) GetAll(ctx context.Context, req models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error) {
	var (
		resp   = models.GetAllAdminsResponse{}
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
		status,
		login,
		password,
        created_at,
        updated_at
        FROM "admin"`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			admin      = models.Admin{}
			Id         sql.NullString
			full_name  sql.NullString
			email      sql.NullString
			age        sql.NullInt16
			status     sql.NullString
			login      sql.NullString
			password   sql.NullString
			created_at sql.NullString
			updateAt   sql.NullString
		)

		if err := rows.Scan(
			&resp.Count,
			&Id,
			&full_name,
			&email,
			&age,
			&status,
			&login,
			&password,
			&created_at,
			&updateAt); err != nil {
			return resp, err
		}
		admin.Updated_at = pkg.NullStringToString(updateAt)
		resp.Admins = append(resp.Admins, models.Admin{
			Id:         Id.String,
			Full_Name:  full_name.String,
			Email:      email.String,
			Age:        uint(age.Int16),
			Status:     status.String,
			Login:      login.String,
			Password:   password.String,
			Created_at: created_at.String,
			Updated_at: updateAt.String,
		})
	}
	return resp, nil
}

func (c *adminRepo) GetByID(ctx context.Context, id string) (models.Admin, error) {
	admin := models.Admin{}
	var (
		full_name  sql.NullString
		email      sql.NullString
		age        sql.NullInt16
		status     sql.NullString
		login      sql.NullString
		password   sql.NullString
		created_at sql.NullString
		updateAt   sql.NullString
	)
	if err := c.db.QueryRow(context.Background(), `select id, full_name, email, age, status, login, password, created_at, updated_at from "admin" where id = $1`, id).Scan(
		&admin.Id,
		&full_name,
		&email,
		&age,
		&status,
		&login,
		&password,
		&created_at,
		&updateAt); err != nil {
		return models.Admin{}, err
	}
	return models.Admin{
		Id:         admin.Id,
		Full_Name:  full_name.String,
		Email:      email.String,
		Age:        uint(age.Int16),
		Status:     status.String,
		Login:      login.String,
		Password:   password.String,
		Created_at: created_at.String,
		Updated_at: updateAt.String,
	}, nil
}

func (c *adminRepo) Delete(ctx context.Context, id string) error {
	query := `delete from "admin" where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
