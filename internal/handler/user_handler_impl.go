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

func (controller *UserHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := dto.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse, err := controller.UserService.Create(request.Context(), userCreateRequest)
	helper.IfPanicError(err)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
