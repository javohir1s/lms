package postgres

import (
	"context"
	"fmt"
	"lms_back/config"
	"lms_back/storage"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	pgPoolConfig.MaxConns = 100
	pgPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(context.Background(), pgPoolConfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	return Store{
		Pool: newPool,
	}, nil
}

func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) Admin() storage.IAdminStorage {
	NewAdmin := NewAdmin(s.Pool)

	return &NewAdmin
}

func (s Store) Branch() storage.IBranchStorage {
	NewBranch := NewBranch(s.Pool)

	return &NewBranch
}

func (s Store) Group() storage.IGroupStorage {
	NewGroup := NewGroup(s.Pool)

	return &NewGroup
}

func (s Store) Lesson() storage.ILessonStorage {
	NewLesson := NewLesson(s.Pool)

	return &NewLesson
}

func (s Store) Payment() storage.IPaymentStorage {
	NewPayment := NewPayment(s.Pool)

	return &NewPayment
}

func (s Store) Schedule() storage.IScheduleStorage {
	NewSchedule := NewSchedule(s.Pool)

	return &NewSchedule
}

func (s Store) Student() storage.IStudentStorage {
	NewStudent := NewStudent(s.Pool)

	return &NewStudent
}

func (s Store) Task() storage.ITaskStorage {
	NewTask := NewTask(s.Pool)

	return &NewTask
}

func (s Store) Teacher() storage.ITeacherStorage {
	NewTeacher := NewTeacher(s.Pool)

	return &NewTeacher
}
