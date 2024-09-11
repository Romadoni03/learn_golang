package handler

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func (handler *UserHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := dto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse, err := handler.UserService.Create(request.Context(), userCreateRequest)
	helper.PanicWithMessage(err, userResponse)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": userResponse},
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (handler *UserHandlerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLogin := dto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userLogin)

	userResponse := handler.UserService.Login(request.Context(), userLogin)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandlerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	message := handler.UserService.Logout(request.Context(), request)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   map[string]string{"message": message},
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *UserHandlerImpl) GetByToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
