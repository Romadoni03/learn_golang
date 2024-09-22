package service

import (
	"context"
	"ecommerce-cloning-app/internal/dto"
	"net/http"
)

type UserService interface {
	Create(ctx context.Context, request dto.UserCreateRequest) (string, error)
	Login(ctx context.Context, request dto.UserCreateRequest) dto.UserLoginResponse
	Logout(ctx context.Context, request *http.Request) string
	GetByToken(ctx context.Context, request *http.Request) dto.UserProfileResponse
	Update(ctx context.Context, request dto.UserUpdateRequest, token string) dto.UserProfileResponse
}
