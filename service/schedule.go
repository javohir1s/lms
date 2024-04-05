package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type scheduleService struct {
	storage storage.IStorage
}

func NewScheduleService(storage storage.IStorage) scheduleService {
	return scheduleService{
		storage: storage,
	}
}

func (u scheduleService) Create(ctx context.Context, schedule models.Schedule) (models.Schedule, error) {

	pKey, err := u.storage.Schedule().Create(ctx, schedule)
	if err != nil {
		fmt.Println("ERROR in service layer while creating schedule", err.Error())
		return models.Schedule{}, err
	}

	return pKey, nil
}

func (u scheduleService) Update(ctx context.Context, schedule models.Schedule) (models.Schedule, error) {

	pKey, err := u.storage.Schedule().Update(ctx, schedule)
	if err != nil {
		fmt.Println("ERROR in service layer while updating schedule", err.Error())
		return models.Schedule{}, err
	}
	return pKey, nil
}

func (u scheduleService) GetByID(ctx context.Context, id string) (models.Schedule, error) {

	pKey, err := u.storage.Schedule().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid schedule", err.Error())
		return models.Schedule{}, err
	}

	return pKey, nil
}

func (u scheduleService) GetAll(ctx context.Context, req models.GetAllSchedulesRequest) (models.GetAllSchedulesResponse, error) {

	pKey, err := u.storage.Schedule().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll schedule", err.Error())
		return models.GetAllSchedulesResponse{}, err
	}

	return pKey, nil
}

func (u scheduleService) Delete(ctx context.Context, id string) error {

	err := u.storage.Schedule().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting schedule", err.Error())
		return err
	}

	return nil
}
