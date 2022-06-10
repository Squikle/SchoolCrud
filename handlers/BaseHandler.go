package handlers

import (
	"DbProj/DAL"
	"errors"
	"net/http"
)

type BaseHandler[T DAL.DbEntity] struct {
	handler Handler[T]
}

func (baseHandler *BaseHandler[T]) Handle(writer http.ResponseWriter, req *http.Request) {
	var err = baseHandler.handle(writer, req)
	if err != nil {
		println(err.Error())
	}
}

func (baseHandler *BaseHandler[T]) handle(writer http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "GET":
		writer.Header().Set("Content-Type", "application/json")
		return (baseHandler.handler).HandleGet(writer, req)
	case "POST":
		return (baseHandler.handler).HandlePost(writer, req)
	case "PUT":
		return (baseHandler.handler).HandlePut(writer, req)
	case "DELETE":
		return (baseHandler.handler).HandleDelete(writer, req)
	}

	return errors.New("method is not supported")
}

func (baseHandler *BaseHandler[T]) HandleGet(_ http.ResponseWriter, _ *http.Request) error {
	return getNotImplementedErr()
}
func (baseHandler *BaseHandler[T]) HandlePost(_ http.ResponseWriter, _ *http.Request) error {
	return getNotImplementedErr()
}
func (baseHandler *BaseHandler[T]) HandlePut(_ http.ResponseWriter, _ *http.Request) error {
	return getNotImplementedErr()
}
func (baseHandler *BaseHandler[T]) HandleDelete(_ http.ResponseWriter, _ *http.Request) error {
	return getNotImplementedErr()
}
func getNotImplementedErr() error {
	return errors.New("method is not implemented")
}
