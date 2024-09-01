package service

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.UserCreateRequest) (string, error) {
	err := service.Validate.Struct(request)
	helper.IfPanicError(err)

	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Id:        helper.GenerateId(),
		NoTelepon: request.NoTelepon,
		Password:  helper.HashingPassword(request.Password),
		Username:  helper.GeneratedUsername(),
	}

	errService := service.UserRepository.Insert(ctx, tx, user)

	if errService != nil {
		return "", errService
	} else {
		return "success create new user", nil
	}

}
