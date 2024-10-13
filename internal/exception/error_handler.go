package exception

import (
	"ecommerce-cloning-app/internal/dto"
	"ecommerce-cloning-app/internal/helper"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundErrors(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	if unauthorizedErrors(writer, request, err) {
		return
	}
	if internalServerErrors(writer, request, err) {
		return
	}
}

func validationErrors(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(ValidationError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   map[string]string{"message": exception.Error},
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func notFoundErrors(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := dto.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   map[string]string{"message": exception.Error},
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func unauthorizedErrors(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(UnauthorizedError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := dto.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   map[string]string{"message": exception.Error},
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerErrors(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(InternalServerError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		webResponse := dto.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   map[string]string{"message": exception.Error},
		}
		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}
