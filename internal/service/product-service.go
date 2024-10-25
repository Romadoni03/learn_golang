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
	"github.com/shopspring/decimal"
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
		Price:             decimal.NewFromInt(request.Price),
		Stock:             request.Stock,
		Wholesaler:        request.Wholesaler,
		ShippingCost:      decimal.NewFromInt(request.ShippingCost),
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

func (service *ProductService) FindAll(ctx context.Context, token string) []dto.ProductRespone {
	tx, errSQL := service.DB.Begin()
	if errSQL != nil {
		logger.Logging().Error(errSQL)
		panic(exception.NewInternalServerError("internal server error"))
	}
	defer helper.CommitOrRollback(tx)

	user, errUser := service.UserRepository.FindFirstByToken(ctx, tx, token)
	if errUser != nil {
		logger.Logging().Error(errUser)
		panic(exception.NewInternalServerError("user by token " + token + "is not found"))
	}
	logger.Logging().Info("Request from Product : " + user.Username + " call FindAll func in ProductService")

	store, errStore := service.StoreRepository.FindByUser(ctx, tx, user)
	if errStore != nil {
		logger.Logging().Error(errStore)
		panic(exception.NewInternalServerError("store by username " + user.Username + "is not found"))
	}

	products := service.ProductRepository.FindAll(ctx, tx, store)
	if products == nil {
		logger.Logging().Warning("product is empty")
		panic(exception.NewNotFoundError("product is empty"))
	}

	var productResponses []dto.ProductRespone
	for _, product := range products {
		productResponses = append(productResponses, dto.ProductRespone{
			Id:         product.Id,
			Name:       product.Name,
			Wholesaler: product.Wholesaler,
			Price:      product.Price,
			Stock:      product.Stock,
		})
	}

	return productResponses
}

func (service *ProductService) FindById(ctx context.Context, productId string) dto.ProductCreateUpdateResponse {
	tx, errSQL := service.DB.Begin()
	if errSQL != nil {
		logger.Logging().Error(errSQL)
		panic(exception.NewInternalServerError("internal server error"))
	}
	defer helper.CommitOrRollback(tx)

	logger.Logging().Info("Request from Product : " + productId + " call FindById func in ProductService")
	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		logger.Logging().Error(err)
		panic(exception.NewNotFoundError("product is not found"))
	}

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

func (service *ProductService) Update(ctx context.Context, request dto.ProductCreateUpdateRequest, productId string, token string) dto.ProductCreateUpdateResponse {
	logger.Logging().Info("Request from Product : " + request.Name + " call Update function")
	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		logger.Logging().Error("Err :" + errValidate.Error() + "field can not be null")
		panic(exception.NewValidationError("field can not be null"))
	}

	tx, errSQL := service.DB.Begin()
	if errSQL != nil {
		logger.Logging().Error(errSQL)
		panic(exception.NewInternalServerError("internal server error"))
	}
	defer helper.CommitOrRollback(tx)
	user, errUser := service.UserRepository.FindFirstByToken(ctx, tx, token)
	if errUser != nil {
		logger.Logging().Error(errUser)
		panic(exception.NewInternalServerError("user by token " + token + "is not found"))
	}
	logger.Logging().Info("Request from Product : " + user.Username + " call FindAll func in ProductService")

	store, errStore := service.StoreRepository.FindByUser(ctx, tx, user)
	if errStore != nil {
		logger.Logging().Error(errStore)
		panic(exception.NewInternalServerError("store by username " + user.Username + "is not found"))
	}

	product, errProduct := service.ProductRepository.FindById(ctx, tx, productId)
	if errProduct != nil {
		logger.Logging().Error("Err :" + errProduct.Error())
		panic(exception.NewNotFoundError("product is not found"))
	}

	product.PhotoProduct = request.PhotoProduct
	product.Name = request.Name
	product.Category = request.Category
	product.Description = request.Description
	product.DangeriousProduct = request.DangeriousProduct
	product.Price = decimal.NewFromInt(request.Price)
	product.Stock = request.Stock
	product.Wholesaler = request.Wholesaler
	product.ShippingCost = decimal.NewFromInt(request.ShippingCost)
	product.ShippingInsurance = request.ShippingInsurance
	product.Conditions = request.Conditions
	product.PreOrder = request.PreOrder
	product.Status = request.Status
	product.LastUpdatedAt = helper.GeneratedTimeNow()

	productResult := service.ProductRepository.Update(ctx, tx, product, store)

	return dto.ProductCreateUpdateResponse{
		Id:                productResult.Id,
		PhotoProduct:      productResult.PhotoProduct,
		Name:              productResult.Name,
		Category:          productResult.Category,
		Description:       productResult.Description,
		DangeriousProduct: productResult.DangeriousProduct,
		Price:             productResult.Price,
		Stock:             productResult.Stock,
		Wholesaler:        productResult.Wholesaler,
		ShippingCost:      productResult.ShippingCost,
		ShippingInsurance: productResult.ShippingInsurance,
		Conditions:        productResult.Conditions,
		PreOrder:          productResult.PreOrder,
		Status:            productResult.Status,
	}

}

func (service *ProductService) Delete(ctx context.Context, productId, token string) string {
	tx, errSQL := service.DB.Begin()
	if errSQL != nil {
		logger.Logging().Error(errSQL)
		panic(exception.NewInternalServerError("internal server error"))
	}
	defer helper.CommitOrRollback(tx)
	user, errUser := service.UserRepository.FindFirstByToken(ctx, tx, token)
	if errUser != nil {
		logger.Logging().Error(errUser)
		panic(exception.NewInternalServerError("user by token " + token + "is not found"))
	}
	logger.Logging().Info("Request from Product : " + user.Username + " call Delete func in ProductService")

	store, errStore := service.StoreRepository.FindByUser(ctx, tx, user)
	if errStore != nil {
		logger.Logging().Error(errStore)
		panic(exception.NewInternalServerError("store by username " + user.Username + "is not found"))
	}

	product, errProduct := service.ProductRepository.FindById(ctx, tx, productId)
	if errProduct != nil {
		logger.Logging().Error("Err :" + errProduct.Error())
		panic(exception.NewNotFoundError("product is not found"))
	}

	logger.Logging().Info("Request from Product : " + productId + " call Delete func in ProductService")
	err := service.ProductRepository.Delete(ctx, tx, product, store)
	if err != nil {
		logger.Logging().Error(err)
		panic(exception.NewNotFoundError("product is not found"))
	} else {
		return "Success delete product :" + product.Name + "with id :" + product.Id
	}
}
