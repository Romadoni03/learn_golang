package handler

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/service"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logger.LogHandler(request).Info("Incoming Request")
	userCreateRequest := dto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := handler.UserService.Create(request.Context(), userCreateRequest)
	logger.LogHandler(request).Info(userResponse)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": userResponse},
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (handler *UserHandler) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logger.LogHandler(request).Info("Incoming Request")
	loginRequest := dto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	userResponse, refreshToken := handler.UserService.Login(request.Context(), loginRequest)
	logger.LogHandler(request).Info(userResponse)
	logger.Logging().Info("access_token :" + userResponse.AccessToken)
	logger.Logging().Info("refresh_token :" + refreshToken)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
	})

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandler) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	message := handler.UserService.Logout(request.Context(), request)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": message},
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
	})

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandler) FindUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := handler.UserService.FindUser(request.Context(), request)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := dto.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userResponse := handler.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
