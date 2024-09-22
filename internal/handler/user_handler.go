package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserHandler interface {
	Create(writer http.ResponseWriter, Request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetByToken(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
