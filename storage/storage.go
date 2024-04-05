package storage

import (
	"context"
	"lms_back/api/models"
)

type IStorage interface {
	CloseDB()
	Admin() IAdminStorage
	Branch() IBranchStorage
	Group() IGroupStorage
	Payment() IPaymentStorage
	Student() IStudentStorage
	Teacher() ITeacherStorage
	Schedule() IScheduleStorage
	Task() ITaskStorage
	Lesson() ILessonStorage
}

type IAdminStorage interface {
	Create(context.Context, models.Admin) (models.Admin, error)
	GetAll(ctx context.Context, request models.GetAllAdminsRequest) (models.GetAllAdminsResponse, error)
	GetByID(ctx context.Context, id string) (models.Admin, error)
	Update(context.Context, models.Admin) (models.Admin, error)
	Delete(context.Context, string) error
}

type IBranchStorage interface {
	Create(context.Context, models.Branch) (models.Branch, error)
	GetAll(ctx context.Context, request models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error)
	GetByID(ctx context.Context, id string) (models.Branch, error)
	Update(context.Context, models.Branch) (models.Branch, error)
	Delete(context.Context, string) error
}

type IGroupStorage interface {
	Create(context.Context, models.Group) (models.Group, error)
	GetAll(ctx context.Context, request models.GetAllGroupsRequest) (models.GetAllGroupsResponse, error)
	GetByID(ctx context.Context, id string) (models.Group, error)
	Update(context.Context, models.Group) (models.Group, error)
	Delete(context.Context, string) error
}

type ILessonStorage interface {
	Create(context.Context, models.Lesson) (models.Lesson, error)
	GetAll(ctx context.Context,req models.GetAllLessonsRequest) (models.GetAllLessonsResponse, error)
	GetByID(ctx context.Context, id string) (models.Lesson, error)
	Update(context.Context, models.Lesson) (models.Lesson, error)
	Delete(context.Context, string) error
}

type IPaymentStorage interface {
	Create(context.Context, models.Payment) (models.Payment, error)
	GetAll(ctx context.Context, request models.GetAllPaymentsRequest) (models.GetAllPaymentsResponse, error)
	GetByID(ctx context.Context, id string) (models.Payment, error)
	Update(context.Context, models.Payment) (models.Payment, error)
	Delete(context.Context, string) error
}

type IStudentStorage interface {
	Create(context.Context,models.Student) (models.Student, error)
	GetAll(ctx context.Context, request models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetByID(ctx context.Context,id string) (models.Student, error)
	Update(context.Context, models.Student) (models.Student, error)
	Delete(context.Context, string) error
}

type ITeacherStorage interface {
	Create(context.Context,models.Teacher) (models.Teacher, error)
	GetAll(ctx context.Context, request models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetByID(ctx context.Context, id string) (models.Teacher, error)
	Update(context.Context, models.Teacher) (models.Teacher, error)
	Delete(context.Context, string) error
}

type IScheduleStorage interface {
	Create(context.Context, models.Schedule) (models.Schedule, error)
	GetAll(ctx context.Context, request models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error)
	GetByID(ctx context.Context, id string) (models.Schedule, error)
	Update(context.Context, models.Schedule) (models.Schedule, error)
	Delete(context.Context, string) error
}

type ITaskStorage interface {
	Create(context.Context, models.Task) (models.Task, error)
	GetAll(ctx context.Context, req models.GetAllTasksRequest) (models.GetAllTasksResponse, error)
	GetByID(ctx context.Context, id string) (models.Task, error)
	Update(context.Context, models.Task) (models.Task, error)
	Delete(context.Context, string) error
}

