package service

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/repository"

	"github.com/go-playground/validator/v10"
)

type StoreServiceImpl struct {
	StoreRepository repository.StoreRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func (service *StoreServiceImpl) Create(ctx context.Context, request dto.StoreCreateRequest) {

}
