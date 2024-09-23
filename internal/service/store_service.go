package service

import (
	"context"
	"ecommerce-cloning-app/internal/dto"
)

type StoreService interface {
	Create(ctx context.Context, request dto.StoreCreateRequest)
}
