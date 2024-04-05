package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"lms_back/api/models"
	"lms_back/pkg"
	"strconv"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepo struct {
	db *pgxpool.Pool
}

func NewGroup(db *pgxpool.Pool) GroupRepo {
	return GroupRepo{
		db: db,
	}
}

func (g *GroupRepo) Create(ctx context.Context, group models.Group) (models.Group, error) {
	id := uuid.New()
	var group_unique_id string

	maxQuery := `SELECT COALESCE(MAX(group_id), 'GR-0000000') FROM "group"`
	err := g.db.QueryRow(context.Background(), maxQuery).Scan(&group_unique_id)
	if err != nil {
		if err.Error() != "can't scan into dest[0]: cannot scan null into *string" && err.Error() != "no rows in result set" {
			return models.Group{}, err
		} else {
			group_unique_id = "GR-0000000"
		}
	}
	digit := 0
	if len(group_unique_id) > 2 {
		digit, err = strconv.Atoi(group_unique_id[3:])
		if err != nil {
			return models.Group{}, err
		}
	}
	query := `INSERT INTO "group" (
		id,
		group_id,
		branch_id,
		teacher_id,
		type,
		created_at) 
		VALUES($1,$2,$3,$4,$5,CURRENT_TIMESTAMP)
		`

	_, err = g.db.Exec(context.Background(), query,
		id.String(),
		"Gr-"+pkg.GetSerialId(digit),
		group.Branch_id,
		group.Teacher_id,
		group.Type)
	if err != nil {
		return models.Group{}, err
	}
	return models.Group{
		Id:         group.Id,
		Group_id:   group.Group_id,
		Branch_id:  group.Branch_id,
		Teacher_id: group.Teacher_id,
		Type:       group.Type,
		Created_at: group.Created_at,
	}, nil
}

func (g *GroupRepo) Update(ctx context.Context, group models.Group) (models.Group, error) {
	query := `UPDATE "group" SET
		type=$1,
		updated_at=CURRENT_TIMESTAMP
		WHERE id=$2`

	_, err := g.db.Exec(context.Background(), query, group.Type, group.Id)
	if err != nil {
		return models.Group{}, nil
	}
	return models.Group{
		Id:         group.Id,
		Group_id:   group.Group_id,
		Branch_id:  group.Branch_id,
		Teacher_id: group.Teacher_id,
		Type:       group.Type,
		Created_at: group.Created_at,
		Updated_at: group.Updated_at,
	}, nil
}

func (g *GroupRepo) GetAll(ctx context.Context, req models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error) {
	var (
		resp   = models.GetAllGroupsResponse{}
		filter = ""
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter += fmt.Sprintf(` and name ILIKE  '%%%v%%' `, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)
	fmt.Println("filter: ", filter)

	rows, err := g.db.Query(context.Background(), `SELECT count (id) OVER(),
        id,
        group_id,
        branch_id,
        teacher_id,
        type,
        created_at,
        updated_at FROM "group"`+filter+``)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			group      = models.Group{}
			group_id   sql.NullString
			branch_id  sql.NullString
			teacher_id sql.NullString
			Type       sql.NullString
			created_at sql.NullString
			updateAt   sql.NullString
		)
		if err := rows.Scan(
			&resp.Count,
			&group.Id,
			&group_id,
			&branch_id,
			&teacher_id,
			&Type,
			&created_at,
			&updateAt); err != nil {
			return resp, err
		}
		group.Updated_at = pkg.NullStringToString(updateAt)
		resp.Groups = append(resp.Groups, models.Group{
			Group_id:   group_id.String,
			Branch_id:  branch_id.String,
			Teacher_id: teacher_id.String,
			Type:       Type.String,
			Created_at: created_at.String,
			Updated_at: updateAt.String,
		})
	}
	return resp, nil
}

func (g *GroupRepo) GetByID(ctx context.Context, id string) (models.Group, error) {
	var (
		group      = models.Group{}
		group_id   sql.NullString
		branch_id  sql.NullString
		teacher_id sql.NullString
		Type       sql.NullString
		created_at sql.NullString
		updateAt   sql.NullString
	)
	if err := g.db.QueryRow(context.Background(), `SELECT id, group_id, branch_id, teacher_id, type, created_at, updated_at FROM "group" WHERE id = $1`, id).Scan(
		&group.Id,
		&group_id,
		&branch_id,
		&teacher_id,
		&Type,
		&created_at,
		&updateAt); err != nil {
		return models.Group{}, err
	}
	return models.Group{
		Group_id:   group_id.String,
		Branch_id:  branch_id.String,
		Teacher_id: teacher_id.String,
		Type:       Type.String,
		Created_at: created_at.String,
		Updated_at: updateAt.String,
	}, nil
}

func (g *GroupRepo) Delete(ctx context.Context, id string) error {
	query := `delete from "group" where id = $1`
	_, err := g.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
