package handler

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"ecommerce-cloning-app/internal/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (handler *ProductHandler) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("API-KEY")
	productRequest := dto.ProductCreateUpdateRequest{}
	helper.ReadFromRequestBody(request, &productRequest)

	result := handler.ProductService.Create(request.Context(), productRequest, token)
	logger.LogHandler(request).Info(result)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *ProductHandler) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("API-KEY")

	productResponses := handler.ProductService.FindAll(request.Context(), token)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
