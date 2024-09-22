package service

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/repository"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
		PhotoProfile:        "account_profile.png",
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

	data, errCheck := service.UserRepository.FindByPhone(ctx, tx, request.NoTelepon)
	if errCheck != nil {
		panic(exception.NewUnauthorizedError("username or password is wrong"))
	}
	errCheckPw := helper.CompiringPassword(data.Password, request.Password)

	if errCheckPw != nil {
		panic(exception.NewUnauthorizedError("username or password is wrong"))
	}
	user := entity.User{
		Username:       data.Username,
		NoTelepon:      data.NoTelepon,
		Token:          uuid.NewString(),
		TokenExpiredAt: time.Now().Local().UnixMilli() + (1000 * 60 * 60 * 24 * 7),
	}
	errToken := service.UserRepository.UpdateToken(ctx, tx, user)
	helper.PanicWithMessage(errToken, "failed set token")

	return dto.UserLoginResponse{
		NoTelepon:      user.NoTelepon,
		Username:       user.Username,
		Token:          user.Token,
		TokenExpiredAt: user.TokenExpiredAt,
	}
}

func (service *UserServiceImpl) Logout(ctx context.Context, request *http.Request) string {
	token := request.Header.Get("API-KEY")
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)
	errLogout := service.UserRepository.DeleteToken(ctx, tx, token)
	if errLogout != nil {
		panic(exception.NewUnauthorizedError("failed to logout"))
	}
	return "success logout"
}

func (service *UserServiceImpl) GetByToken(ctx context.Context, request *http.Request) dto.UserProfileResponse {
	token := request.Header.Get("API-KEY")
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	var nameStore string
	// var userProfile string

	user, err := service.UserRepository.GetByToken(ctx, tx, token)
	if err != nil {
		panic(exception.NewNotFoundError("user not found"))
	}
	if !user.Store.Name.Valid {
		nameStore = ""
	} else {
		nameStore = user.Store.Name.String
	}

	return dto.UserProfileResponse{
		Username:     user.Username,
		Name:         user.Name,
		Email:        user.Email,
		NoTelepon:    user.NoTelepon,
		PhotoProfile: helper.GetImage(user.PhotoProfile),
		NameStore:    nameStore,
		Gender:       user.Gender,
		BirthDate:    user.BirthDate,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request dto.UserUpdateRequest, token string) dto.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	user, errGetByToken := service.UserRepository.GetByToken(ctx, tx, token)
	if errGetByToken != nil {
		panic(exception.NewNotFoundError("User by token not found"))
	}
	//Usernames can only be updated once every 30 days
	userLastUsername := user.LastUpdatedUsername.UnixMilli()

	if (userLastUsername + (1000 * 60 * 60 * 24 * 30)) <= time.Now().Local().UnixMilli() {
		panic(exception.NewUnauthorizedError("can't update username before 30 days"))
	}

	if user.Store.Name.String == "" {
		request.NameStore = ""
	}
	// set user
	user.Username = request.Username
	user.Name = request.Name
	user.Store.Name.String = request.NameStore
	user.Gender = request.Gender
	user.BirthDate = request.BirthDate
	user.LastUpdatedUsername = helper.GeneratedTimeNow()
	errUpdate := service.UserRepository.Update(ctx, tx, user)
	if errUpdate != nil {
		panic(exception.NewUnauthorizedError(errUpdate.Error()))
	}

	return dto.UserProfileResponse{
		Username:     user.Username,
		Name:         user.Name,
		Email:        user.Email,
		NoTelepon:    user.NoTelepon,
		PhotoProfile: helper.GetImage(user.PhotoProfile),
		NameStore:    user.Store.Name.String,
		Gender:       user.Gender,
		BirthDate:    user.BirthDate,
	}
}
