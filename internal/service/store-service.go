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

func (service *StoreService) Create(ctx context.Context, request dto.StoreCreateRequest, token string) dto.StoreCreateResponse {
	logger.Logging().Info("Request from Store : " + request.Name + " call Create func in StoreService")
	errValidate := service.Validate.Struct(request)
	helper.IfPanicError(errValidate)

	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, errCheckToken := service.UserRepository.FindFirstByToken(ctx, tx, token)
	if errCheckToken != nil {
		logger.Logging().Error("Err :" + errCheckToken.Error() + "user whith token " + token + " is not found")
		panic(exception.NewNotFoundError("user whith token " + token + " is not found"))
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

func (service *StoreService) Delete(ctx context.Context, token string) dto.StoreCreateResponse {
	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, _ := service.UserRepository.FindFirstByToken(ctx, tx, token)
	logger.Logging().Info("Request from Store : " + user.Username + " call Delete func in StoreService")

	storeResult, _ := service.StoreRepository.FindByUser(ctx, tx, user)
	err := service.StoreRepository.Delete(ctx, tx, storeResult)
	helper.IfPanicError(err)

	return dto.StoreCreateResponse{
		Name:    storeResult.Name.String,
		Message: "Success Delete Store",
	}

}

func (service *StoreService) FindByUser(ctx context.Context, token string) dto.StoreGetResponse {
	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, _ := service.UserRepository.FindFirstByToken(ctx, tx, token)
	logger.Logging().Info("Request from Store : " + user.Username + " call FindByUser func in StoreService")

	storeResult, err := service.StoreRepository.FindByUser(ctx, tx, user)
	helper.IfPanicError(err)

	return dto.StoreGetResponse{
		Id:          storeResult.StoreId,
		Name:        storeResult.Name.String,
		Logo:        storeResult.Logo,
		Description: storeResult.Description,
		NoTelepon:   storeResult.NoTelepon,
		Email:       user.Email,
	}
}
