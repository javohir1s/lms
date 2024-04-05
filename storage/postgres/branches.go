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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranch(db *pgxpool.Pool) branchRepo {
	return branchRepo{
		db: db,
	}
}

func (c *branchRepo) Create(ctx context.Context, branch models.Branch) (models.Branch, error) {

	id := uuid.New()
	query := `INSERT INTO branches (
		id,
		name,
		address,
		created_at)
		VALUES($1,$2,$3,CURRENT_TIMESTAMP) 
	`

	_, err := c.db.Exec(context.Background(), query,
		id.String(),
		branch.Name,
		branch.Address)

	if err != nil {
		return models.Branch{}, err
	}
	return models.Branch{
		Id:        branch.Id,
		Name:      branch.Name,
		Address:   branch.Address,
		CreatedAt: branch.CreatedAt,
	}, nil
}

func (c *branchRepo) Update(ctx context.Context, branch models.Branch) (models.Branch, error) {
	query := `update branches set 
	name=$1,
	address=$2,
	updated_at=CURRENT_TIMESTAMP
	WHERE id = $3 AND deleted_at = 0
	`
	_, err := c.db.Exec(context.Background(), query,
		branch.Name,
		branch.Address,
		branch.Id,
	)
	if err != nil {
		return models.Branch{}, err
	}
	return models.Branch{
		Id:        branch.Id,
		Name:      branch.Name,
		Address:   branch.Address,
		CreatedAt: branch.CreatedAt,
		UpdatedAt: branch.UpdatedAt,
	}, nil
}

func (c *branchRepo) GetAll(ctx context.Context, req models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {
	var (
		resp   = models.GetAllBranchesResponse{}
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
        name,
        address,
        created_at,
        updated_at,
        deleted_at FROM branches WHERE deleted_at = 0`+filter+``)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			branch     = models.Branch{}
			id         sql.NullString
			name       sql.NullString
			address    sql.NullString
			created_at sql.NullString
			updateAt   sql.NullString
			deleted_at sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&created_at,
			&updateAt,
			&deleted_at); err != nil {
			return resp, err
		}
		branch.UpdatedAt = pkg.NullStringToString(updateAt)
		resp.Branches = append(resp.Branches, models.Branch{
			Id:        id.String,
			Name:      name.String,
			Address:   address.String,
			CreatedAt: created_at.String,
			UpdatedAt: updateAt.String,
			DeletedAt: deleted_at.String,
		})
	}
	return resp, nil
}

func (c *branchRepo) GetByID(ctx context.Context, id string) (models.Branch, error) {

	var (
		branch     = models.Branch{}
		name       sql.NullString
		address    sql.NullString
		created_at sql.NullString
		updateAt   sql.NullString
		deleted_at sql.NullString
	)

	if err := c.db.QueryRow(context.Background(), `select id, name, address, created_at, updated_at, deleted_at from branches where id = $1`, id).Scan(
		&branch.Id,
		&name,
		&address,
		&created_at,
		&updateAt,
		&deleted_at); err != nil {
		return models.Branch{}, err
	}
	return models.Branch{
		Id:        branch.Id,
		Name:      name.String,
		Address:   address.String,
		CreatedAt: created_at.String,
		UpdatedAt: updateAt.String,
		DeletedAt: deleted_at.String,
	}, nil
}

func (c *branchRepo) Delete(ctx context.Context, id string) error {
	query := `delete from branches where id = $1`
	_, err := c.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
