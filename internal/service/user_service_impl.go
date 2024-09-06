package service

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/repository"
	"errors"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.UserCreateRequest) (string, error) {
	err := service.Validate.Struct(request)
	helper.PanicWithMessage(err, "No Telepon or Password can not be null")

	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	_, errCheck := service.UserRepository.FindByPhone(ctx, tx, request.NoTelepon)

	if errCheck == nil {
		return "No telepon is already", errors.New("user is already")
	}

	user := entity.User{
		NoTelepon:           request.NoTelepon,
		Password:            helper.HashingPassword(request.Password),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        helper.EncodeImageName("account_profile.png"),
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
		TokenExpiredAt:      0,
	}

	errService := service.UserRepository.Insert(ctx, tx, user)

	if errService != nil {
		return "failed to create new user", errService
	} else {
		return "success create new user", nil
	}

}

func (service *UserServiceImpl) Login(ctx context.Context, request dto.UserCreateRequest) dto.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicWithMessage(err, "No Telepon and password can not be null")

	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	user, errCheck := service.UserRepository.FindByPhone(ctx, tx, request.NoTelepon)
	helper.PanicWithMessage(errCheck, "user not found")
	errCheckPw := helper.CompiringPassword(user.Password, request.Password)

	helper.PanicWithMessage(errCheckPw, "password wrong")
	return dto.UserLoginResponse{
		NoTelepon: user.NoTelepon,
		Username:  user.Username,
		Name:      user.Name,
	}
}
