package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type lessonService struct {
	storage storage.IStorage
}

func NewLessonService(storage storage.IStorage) lessonService {
	return lessonService{
		storage: storage,
	}
}

func (u lessonService) Create(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {

	pKey, err := u.storage.Lesson().Create(ctx, lesson)
	if err != nil {
		fmt.Println("ERROR in service layer while creating lesson", err.Error())
		return models.Lesson{}, err
	}

	return pKey, nil
}

func (u lessonService) Update(ctx context.Context, lesson models.Lesson) (models.Lesson, error) {

	pKey, err := u.storage.Lesson().Update(ctx, lesson)
	if err != nil {
		fmt.Println("ERROR in service layer while updating lesson", err.Error())
		return models.Lesson{}, err
	}

	return pKey, nil
}

func (u lessonService) GetByID(ctx context.Context, id string) (models.Lesson, error) {

	pKey, err := u.storage.Lesson().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid lesson", err.Error())
		return models.Lesson{}, err
	}

	return pKey, nil
}

func (u lessonService) GetAll(ctx context.Context, req models.GetAllLessonsRequest) (models.GetAllLessonsResponse, error) {

	pKey, err := u.storage.Lesson().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll lesson", err.Error())
		return models.GetAllLessonsResponse{}, err
	}

	return pKey, nil
}

func (u lessonService) Delete(ctx context.Context, id string) error {

	err := u.storage.Lesson().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting lesson", err.Error())
		return err
	}

	return nil
}
