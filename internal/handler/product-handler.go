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

	productResponse := handler.ProductService.Create(request.Context(), productRequest, token)
	logger.LogHandler(request).Info(productResponse)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
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

func (handler *ProductHandler) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	productResponse := handler.ProductService.FindById(request.Context(), productId)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (handler *ProductHandler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	productRequest := dto.ProductCreateUpdateRequest{}
	helper.ReadFromRequestBody(request, &productRequest)

	productResponse := handler.ProductService.Update(request.Context(), productRequest, productId)
	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
