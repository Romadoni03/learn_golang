package handler

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StoreHandler struct {
	StoreService *service.StoreService
}

func (handler *StoreHandler) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	storeCreateRequest := dto.StoreCreateRequest{}
	helper.ReadFromRequestBody(request, &storeCreateRequest)

	result := handler.StoreService.Create(request.Context(), storeCreateRequest)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *StoreHandler) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := handler.StoreService.Delete(request.Context())

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *StoreHandler) FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	result := handler.StoreService.FindByUser(request.Context())

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *StoreHandler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	storeUpdateRequest := dto.StoreUpdateRequest{}
	helper.ReadFromRequestBody(request, &storeUpdateRequest)

	result := handler.StoreService.Update(request.Context(), storeUpdateRequest)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
