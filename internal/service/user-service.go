package service

import (
	"context"
	"database/sql"
	entity "ecommerce-cloning-app/entities"
	"ecommerce-cloning-app/internal/auth"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/repository"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	UserRepository *repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *UserService) Create(ctx context.Context, request dto.UserCreateRequest) string {
	logger.Logging().Info("request from phone : " + request.NoTelepon + " call Create Func In Service")
	err := service.Validate.Struct(request)
	if err != nil {
		logger.Logging().Error("Err :" + err.Error() + "username or password can not be null")
		panic(exception.NewValidationError("username or password can not be null"))
	}

	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	_, errCheck := service.UserRepository.FindByPhone(ctx, tx, request.NoTelepon)

	if errCheck == nil {
		logger.Logging().Error("user is already")
		panic(exception.NewInternalServerError("user is already"))
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
		logger.Logging().Error(errService)
		panic(exception.NewInternalServerError("failed to create new user"))
	}

	return "success create new user"

}

func (service *UserService) Login(ctx context.Context, request dto.UserCreateRequest) (dto.UserLoginResponse, string) {
	logger.Logging().Info("request from phone : " + request.NoTelepon + " call Login Func In Service")
	err := service.Validate.Struct(request)
	if err != nil {
		logger.Logging().Error("Err :" + err.Error() + "username or password can not be null")
		panic(exception.NewValidationError("username or password can not be null"))
	}

	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	data, errCheck := service.UserRepository.FindByPhone(ctx, tx, request.NoTelepon)
	if errCheck != nil {
		logger.Logging().Error("username or password is wrong")
		panic(exception.NewUnauthorizedError("username or password is wrong"))
	}

	errCheckPw := helper.CompiringPassword(data.Password, request.Password)
	if errCheckPw != nil {
		logger.Logging().Error("username or password is wrong")
		panic(exception.NewUnauthorizedError("username or password is wrong"))
	}

	user := entity.User{
		Username:  data.Username,
		NoTelepon: data.NoTelepon,
	}
	tokenJWT, errToken := auth.GenerateJWT(user.NoTelepon)
	if errToken != nil {
		logger.Logging().Error("Failed set Token")
		panic(exception.NewUnauthorizedError("failed set token"))
	}

	return dto.UserLoginResponse{
		Message:   "Login Success",
		NoTelepon: user.NoTelepon,
		Username:  user.Username,
	}, tokenJWT
}

func (service *UserService) Logout(ctx context.Context, request *http.Request) string {
	token := request.Header.Get("API-KEY")
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	errLogout := service.UserRepository.DeleteToken(ctx, tx, token)
	if errLogout != nil {
		logger.Logging().Error("Err : " + errLogout.Error() + "failed to logout")
		panic(exception.NewUnauthorizedError("failed to logout"))
	}
	return "success logout"
}

func (service *UserService) GetByToken(ctx context.Context, request *http.Request) dto.UserProfileResponse {
	token := request.Header.Get("API-KEY")
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	var nameStore string
	// var userProfile string

	user, err := service.UserRepository.GetByToken(ctx, tx, token)
	if err != nil {
		logger.Logging().Error("Err : " + err.Error() + "user not found")
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

func (service *UserService) Update(ctx context.Context, request dto.UserUpdateRequest, token string) dto.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	user, errGetByToken := service.UserRepository.GetByToken(ctx, tx, token)
	if errGetByToken != nil {
		logger.Logging().Error("Err : " + errGetByToken.Error() + "User by token not found")
		panic(exception.NewNotFoundError("User by token not found"))
	}
	//Usernames can only be updated once every 30 days
	userLastUsername := user.LastUpdatedUsername.UnixMilli()

	if (userLastUsername + (1000 * 60 * 60 * 24 * 30)) > time.Now().Local().UnixMilli() {
		logger.Logging().Error("can't update username before 30 days")
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
		logger.Logging().Error("failed update user")
		panic(exception.NewUnauthorizedError("failed update user"))
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
