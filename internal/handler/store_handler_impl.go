package handler

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StoreHandlerImpl struct {
	StoreService service.StoreService
}

func (handler *StoreHandlerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("API-KEY")
	storeCreateRequest := dto.StoreCreateRequest{}
	helper.ReadFromRequestBody(request, &storeCreateRequest)

	result := handler.StoreService.Create(request.Context(), storeCreateRequest, token)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *StoreHandlerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("API-KEY")

	result := handler.StoreService.Delete(request.Context(), token)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
