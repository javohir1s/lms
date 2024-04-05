package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type taskService struct {
	storage storage.IStorage
}

func NewTaskService(storage storage.IStorage) taskService {
	return taskService{
		storage: storage,
	}
}

func (u taskService) Create(ctx context.Context, task models.Task) (models.Task, error) {

	pKey, err := u.storage.Task().Create(ctx, task)
	if err != nil {
		fmt.Println("ERROR in service layer while creating task", err.Error())
		return models.Task{}, err
	}

	return pKey, nil
}

func (u taskService) Update(ctx context.Context, task models.Task) (models.Task, error) {

	pKey, err := u.storage.Task().Update(ctx, task)
	if err != nil {
		fmt.Println("ERROR in service layer while updating task", err.Error())
		return models.Task{}, err
	}
	return pKey, nil
}

func (u taskService) GetByID(ctx context.Context, id string) (models.Task, error) {

	pKey, err := u.storage.Task().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid task", err.Error())
		return models.Task{}, err
	}

	return pKey, nil
}

func (u taskService) GetAll(ctx context.Context, req models.GetAllTasksRequest) (models.GetAllTasksResponse, error) {

	pKey, err := u.storage.Task().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll task", err.Error())
		return models.GetAllTasksResponse{}, err
	}

	return pKey, nil
}

func (u taskService) Delete(ctx context.Context, id string) error {

	err := u.storage.Task().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting task", err.Error())
		return err
	}

	return nil
}
