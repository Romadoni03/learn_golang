package service

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
	StoreRepository   *repository.StoreRepository
	UserRepository    *repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func (service *ProductService) Create(ctx context.Context, request dto.ProductCreateUpdateRequest, token string) dto.ProductCreateUpdateResponse {
	logger.Logging().Info("Request from Product : " + request.Name + " call Create function")
	errValidate := service.Validate.Struct(request)
	helper.IfPanicError(errValidate)

	tx, errSQL := service.DB.Begin()
	helper.IfPanicError(errSQL)
	defer helper.CommitOrRollback(tx)

	user, errUser := service.UserRepository.FindFirstByToken(ctx, tx, token)
	if errUser != nil {
		logger.Logging().Error("user with token " + token + "not found")
		panic(exception.NewNotFoundError("user not found"))
	}
	store, errStore := service.StoreRepository.FindByUser(ctx, tx, user)
	if errStore != nil {
		logger.Logging().Error("store with username " + user.Username + "not found")
		panic(exception.NewNotFoundError("store not found"))
	}

	product := entity.Product{
		Id:                uuid.NewString(),
		StoreId:           store.StoreId,
		PhotoProduct:      request.PhotoProduct,
		Name:              request.Name,
		Category:          request.Category,
		Description:       request.Description,
		DangeriousProduct: request.DangeriousProduct,
		Price:             request.Price,
		Stock:             request.Stock,
		Wholesaler:        request.Wholesaler,
		ShippingCost:      request.ShippingCost,
		ShippingInsurance: request.ShippingInsurance,
		Conditions:        request.Conditions,
		PreOrder:          request.PreOrder,
		Status:            request.Status,
		CreatedAt:         helper.GeneratedTimeNow(),
		LastUpdatedAt:     helper.GeneratedTimeNow(),
	}

	errProduct := service.ProductRepository.Insert(ctx, tx, product)
	helper.PanicWithMessage(errProduct, "failed to add product")

	return dto.ProductCreateUpdateResponse{
		Id:                product.Id,
		PhotoProduct:      product.PhotoProduct,
		Name:              product.Name,
		Category:          product.Category,
		Description:       product.Description,
		DangeriousProduct: product.DangeriousProduct,
		Price:             product.Price,
		Stock:             product.Stock,
		Wholesaler:        product.Wholesaler,
		ShippingCost:      product.ShippingCost,
		ShippingInsurance: product.ShippingInsurance,
		Conditions:        product.Conditions,
		PreOrder:          product.PreOrder,
		Status:            product.Status,
	}
}
