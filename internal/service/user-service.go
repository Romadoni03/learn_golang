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
		LastUpdatedUsername: time.Now().Local().Add(time.Hour * 720),
		Name:                "",
		Email:               "",
		PhotoProfile:        "account_profile.png",
		Bio:                 "",
		Gender:              "",
		StatusMember:        "Basic",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
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

	if data.Token != "" {
		logger.Logging().Error(data.NoTelepon + " already login")
		panic(exception.NewUnauthorizedError(data.NoTelepon + " already login"))
	}

	errCheckPw := helper.CompiringPassword(data.Password, request.Password)
	if errCheckPw != nil {
		logger.Logging().Error("username or password is wrong")
		panic(exception.NewUnauthorizedError("username or password is wrong"))
	}

	user := entity.User{
		Username:       data.Username,
		NoTelepon:      data.NoTelepon,
		Token:          auth.GenerateRefreshToken(),
		TokenExpiredAt: time.Now().Local().Add(time.Hour * 24),
	}

	tokenJWT, errToken := auth.GenerateJWT(user.NoTelepon)
	if errToken != nil {
		logger.Logging().Error("Failed set Token")
		panic(exception.NewUnauthorizedError("failed set token"))
	}

	errUpdateToken := service.UserRepository.UpdateToken(ctx, tx, user)
	if errUpdateToken != nil {
		logger.Logging().Error("Failed set Token")
		panic(exception.NewUnauthorizedError("failed set token"))
	}

	return dto.UserLoginResponse{
		Message:     "Login Success",
		NoTelepon:   user.NoTelepon,
		Username:    user.Username,
		AccessToken: tokenJWT,
	}, user.Token
}

func (service *UserService) Logout(ctx context.Context, request *http.Request) string {
	refreshToken := request.Cookies()
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	errLogout := service.UserRepository.DeleteToken(ctx, tx, refreshToken[0].Value)
	if errLogout != nil {
		logger.Logging().Error("Err : " + errLogout.Error() + "failed to logout")
		panic(exception.NewUnauthorizedError("failed to logout"))
	}
	return "success logout"
}

func (service *UserService) FindUser(ctx context.Context, request *http.Request) dto.UserProfileResponse {
	phone := ctx.Value("phone").(string)
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindUser(ctx, tx, phone)
	if err != nil {
		logger.Logging().Error("Err : " + err.Error() + "user not found")
		panic(exception.NewNotFoundError("user not found"))
	}

	logger.Logging().Info("photoname:" + user.PhotoProfile)
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

func (service *UserService) Update(ctx context.Context, request dto.UserUpdateRequest) dto.UserProfileResponse {
	phone := ctx.Value("phone").(string)
	tx, err := service.DB.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)

	user, errGetByToken := service.UserRepository.FindUser(ctx, tx, phone)
	if errGetByToken != nil {
		logger.Logging().Error("Err : " + errGetByToken.Error() + "User not found")
		panic(exception.NewNotFoundError("User by token not found"))
	}

	if request.Username != "" {
		if time.Now().Local().Before(user.LastUpdatedUsername) {
			logger.Logging().Error("can't updated username berfore 30 days")
			panic(exception.NewInternalServerError("can't updated username berfore 30 days"))
		}
		user.Username = request.Username
	}

	photo, errPhoto := helper.UploadPhotoProfile(request.PhotoProfile)
	if errPhoto != nil {
		logger.Logging().Error("Err : " + errPhoto.Error())
		panic(exception.NewNotFoundError(errPhoto.Error()))
	}

	if !user.Store.Name.Valid {
		request.NameStore = ""
	}
	// set user
	user.Name = request.Name
	user.PhotoProfile = photo
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
