package handlers

import (
	"DbProj/DAL"
	"net/http"
)

type Handler[T DAL.DbEntity] interface {
	Handle(writer http.ResponseWriter, req *http.Request)
	HandleGet(writer http.ResponseWriter, req *http.Request) error
	HandlePost(writer http.ResponseWriter, req *http.Request) error
	HandlePut(writer http.ResponseWriter, req *http.Request) error
	HandleDelete(writer http.ResponseWriter, req *http.Request) error
}
