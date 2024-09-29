package service

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type StoreServiceImpl struct {
	StoreRepository repository.StoreRepository
	UserRepository  repository.UserRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func (service *StoreServiceImpl) Create(ctx context.Context, request dto.StoreCreateRequest, token string) dto.StoreCreateResponse {
	logger.Logging().Info("Request from Store : " + request.Name + " call Create func in StoreService")
	errValidate := service.Validate.Struct(request)
	helper.IfPanicError(errValidate)

	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, _ := service.UserRepository.FindFirstByToken(ctx, tx, token)

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

	err := service.StoreRepository.Insert(ctx, tx, store)
	helper.PanicWithMessage(err, "failed to create new store")

	logger.Logging().Info("Success create new store : " + store.Name.String)
	return dto.StoreCreateResponse{Name: store.Name.String, Message: "Success create new store"}
}

func (service *StoreServiceImpl) Delete(ctx context.Context, token string) dto.StoreCreateResponse {
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
