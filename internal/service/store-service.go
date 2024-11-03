package service

import (
	"context"
	"database/sql"
	entity "ecommerce-cloning-app/entities"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type StoreService struct {
	StoreRepository *repository.StoreRepository
	UserRepository  *repository.UserRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func (service *StoreService) Create(ctx context.Context, request dto.StoreCreateRequest) dto.StoreCreateResponse {
	logger.Logging().Info("Request from Store : " + request.Name + " call Create func in StoreService")
	phone := ctx.Value("phone").(string)
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		logger.Logging().Error("Err :" + errValidate.Error())
		panic(exception.NewValidationError(errValidate.Error()))
	}

	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, errCheckToken := service.UserRepository.FindUser(ctx, tx, phone)
	if errCheckToken != nil {
		logger.Logging().Error("Err :" + errCheckToken.Error() + "user is not found")
		panic(exception.NewNotFoundError("user is not found"))
	}

	store := entity.Store{
		StoreId:         uuid.NewString(),
		NoTelepon:       user.NoTelepon,
		Name:            helper.NewNullString(request.Name),
		LastUpdatedName: helper.GeneratedTimeNow(),
		Logo:            "store_logo_default.png",
		Description:     "",
		Status:          "",
		LinkStore:       "",
		TotalComment:    0,
		TotalFollowing:  0,
		TotalFollower:   0,
		TotalProduct:    0,
		Condition:       "",
		CreatedAt:       helper.GeneratedTimeNow(),
	}

	errService := service.StoreRepository.Insert(ctx, tx, store)
	if errService != nil {
		logger.Logging().Error(errService)
		panic(exception.NewInternalServerError("failed to create new store"))
	}

	logger.Logging().Info("Success create new store : " + store.Name.String)
	return dto.StoreCreateResponse{Name: store.Name.String, Message: "Success create new store"}
}

func (service *StoreService) Delete(ctx context.Context) dto.StoreCreateResponse {
	phone := ctx.Value("phone").(string)
	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, errFindUser := service.UserRepository.FindUser(ctx, tx, phone)
	if errFindUser != nil {
		logger.Logging().Error(errFindUser)
		panic(exception.NewInternalServerError("user not found"))
	}
	logger.Logging().Info("Request from Store : " + user.Username + " call Delete func in StoreService")

	storeResult, errFindStore := service.StoreRepository.FindByUser(ctx, tx, user)
	if errFindStore != nil {
		logger.Logging().Error(errFindStore)
		panic(exception.NewInternalServerError("store not found"))
	}
	err := service.StoreRepository.Delete(ctx, tx, storeResult)
	if err != nil {
		logger.Logging().Error(err)
		panic(exception.NewInternalServerError(err.Error()))
	}

	return dto.StoreCreateResponse{
		Name:    storeResult.Name.String,
		Message: "Success Delete Store",
	}

}

func (service *StoreService) FindByUser(ctx context.Context) dto.StoreGetResponse {
	phone := ctx.Value("phone").(string)
	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, errFindUser := service.UserRepository.FindUser(ctx, tx, phone)
	if errFindUser != nil {
		logger.Logging().Error(errFindUser)
		panic(exception.NewInternalServerError("user not found"))
	}
	logger.Logging().Info("Request from Store : " + user.Username + " call FindByUser func in StoreService")

	storeResult, err := service.StoreRepository.FindByUser(ctx, tx, user)
	if err != nil {
		logger.Logging().Error(err)
		panic(exception.NewInternalServerError("store not found"))
	}

	return dto.StoreGetResponse{
		Id:          storeResult.StoreId,
		Name:        storeResult.Name.String,
		Logo:        storeResult.Logo,
		Description: storeResult.Description,
		NoTelepon:   storeResult.NoTelepon,
		Email:       user.Email,
	}
}

func (service *StoreService) Update(ctx context.Context, request dto.StoreUpdateRequest) dto.StoreGetResponse {
	logger.Logging().Info("Request from Store : " + request.Name + " call Update func in StoreService")
	phone := ctx.Value("phone").(string)
	errValidate := service.Validate.Struct(request)
	helper.IfPanicError(errValidate)

	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, errFindUser := service.UserRepository.FindUser(ctx, tx, phone)
	if errFindUser != nil {
		logger.Logging().Error(errFindUser)
		panic(exception.NewInternalServerError("user not found"))
	}
	logger.Logging().Info("Request from Store : " + user.Username + " call Update func in StoreService")

	_, err := service.StoreRepository.FindByUser(ctx, tx, user)
	if err != nil {
		logger.Logging().Error(err)
		panic(exception.NewInternalServerError("store not found"))
	}

	return dto.StoreGetResponse{}
}
