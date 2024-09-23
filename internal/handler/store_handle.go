package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type StoreHandler interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
