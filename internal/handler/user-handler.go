package handler

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/service"
	"net/http"

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

	userResponse := handler.UserService.Login(request.Context(), loginRequest)
	logger.LogHandler(request).Info(userResponse)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandler) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	message := handler.UserService.Logout(request.Context(), request)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": message},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandler) GetByToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := handler.UserService.GetByToken(request.Context(), request)
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

	userResponse := handler.UserService.Update(request.Context(), userUpdateRequest, request.Header.Get("API-KEY"))
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
