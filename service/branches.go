package service

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"lms_back/storage"
)

type branchService struct {
	storage storage.IStorage
}

func NewBranchService(storage storage.IStorage) branchService {
	return branchService{
		storage: storage,
	}
}

func (u branchService) Create(ctx context.Context, branch models.Branch) (models.Branch, error) {

	pKey, err := u.storage.Branch().Create(ctx, branch)
	if err != nil {
		fmt.Println("ERROR in service layer while creating car", err.Error())
		return models.Branch{}, err
	}

	return pKey, nil
}

func (u branchService) Update(ctx context.Context, branch models.Branch) (models.Branch, error) {

	pKey, err := u.storage.Branch().Update(ctx, branch)
	if err != nil {
		fmt.Println("ERROR in service layer while updating branch", err.Error())
		return models.Branch{}, err
	}

	return pKey, nil
}

func (u branchService) GetByID(ctx context.Context, id string) (models.Branch, error) {

	pKey, err := u.storage.Branch().GetByID(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while getbyid branch", err.Error())
		return models.Branch{}, err
	}

	return pKey, nil
}

func (u branchService) GetAll(ctx context.Context, req models.GetAllBranchesRequest) (models.GetAllBranchesResponse, error) {

	pKey, err := u.storage.Branch().GetAll(ctx, req)
	if err != nil {
		fmt.Println("ERROR in service layer while GetAll branch", err.Error())
		return models.GetAllBranchesResponse{}, err
	}

	return pKey, nil
}

func (u branchService) Delete(ctx context.Context, id string) error {

	err := u.storage.Branch().Delete(ctx, id)
	if err != nil {
		fmt.Println("ERROR in service layer while deleting branch", err.Error())
		return err
	}

	return nil
}
