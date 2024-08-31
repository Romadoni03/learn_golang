package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(writer http.ResponseWriter, Request *http.Request, params httprouter.Params)
}
