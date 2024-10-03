package service

import (
	"context"
	"ecommerce-cloning-app/internal/dto"
)

type StoreService interface {
	Create(ctx context.Context, request dto.StoreCreateRequest, token string) dto.StoreCreateResponse
	Delete(ctx context.Context, token string) dto.StoreCreateResponse
	FindByUser(ctx context.Context, token string) dto.StoreGetResponse
}
