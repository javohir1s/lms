package service

import "lms_back/storage"

type IServiceManager interface {
	Admin() adminService
	Branch() branchService
	Group() groupService
	Lesson() lessonService
	Payment() paymentService
	Schedule() scheduleService
	Student() studentService
	Task() taskService
	Teacher() teacherService
}

type Service struct {
	adminService    adminService
	branchService   branchService
	groupService    groupService
	lessonService   lessonService
	paymentService  paymentService
	scheduleService scheduleService
	studentService  studentService
	taskService     taskService
	teacherService  teacherService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.adminService = NewAdminService(storage)
	services.branchService = NewBranchService(storage)
	services.groupService = NewGroupService(storage)
	services.paymentService = NewPaymentService(storage)
	services.scheduleService = NewScheduleService(storage)
	services.studentService = NewStudentService(storage)
	services.taskService = NewTaskService(storage)
	services.lessonService = NewLessonService(storage)
	services.teacherService = NewTeacherService(storage)
	return services
}

func (s Service) Admin() adminService {
	return s.adminService
}

func (s Service) Branch() branchService {
	return s.branchService
}

func (s Service) Group() groupService {
	return s.groupService
}

func (s Service) Lesson() lessonService {
	return s.lessonService
}

func (s Service) Payment() paymentService {
	return s.paymentService
}

func (s Service) Schedule() scheduleService {
	return s.scheduleService
}

func (s Service) Student() studentService {
	return s.studentService
}

func (s Service) Task() taskService {
	return s.taskService
}

func (s Service) Teacher() teacherService {
	return s.teacherService
}
