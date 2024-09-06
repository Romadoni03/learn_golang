package service

import (
	"context"
	"ecommerce-cloning-app/internal/dto"
)

type UserService interface {
	Create(ctx context.Context, request dto.UserCreateRequest) (string, error)
	Login(ctx context.Context, request dto.UserCreateRequest) dto.UserLoginResponse
}
