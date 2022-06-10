package handlers

import (
	"DbProj/DAL"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	_ "net/http"
	"reflect"
	"strings"
)

type DefaultHandler[T DAL.DbEntity] struct {
	baseHandler *BaseHandler[T]
	Db          *sql.DB
}

func NewDefaultHandler[T DAL.DbEntity](db *sql.DB) *DefaultHandler[T] {
	h := DefaultHandler[T]{Db: db}
	h.baseHandler = &BaseHandler[T]{handler: &h}
	return &h
}

func (defaultHandler *DefaultHandler[T]) Handle(writer http.ResponseWriter, req *http.Request) {
	defaultHandler.baseHandler.Handle(writer, req)
}

func (defaultHandler *DefaultHandler[T]) HandleGet(writer http.ResponseWriter, _ *http.Request) error {
	var db = sqlx.NewDb(defaultHandler.Db, "mysql")
	result := []T{}

	var entityName = GetTypeName[T]() + "s"
	var query = "SELECT * FROM %s;"
	query = fmt.Sprintf(query, entityName)

	err := db.Select(&result, query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return err
	}
	err = json.NewEncoder(writer).Encode(result)
	return err
}

func (defaultHandler *DefaultHandler[T]) HandlePost(writer http.ResponseWriter, req *http.Request) error {
	var db = sqlx.NewDb(defaultHandler.Db, "mysql")
	var t T
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		return err
	}
	var entityType = reflect.TypeOf(t).Elem()
	var fieldsCount = entityType.NumField()

	vars := mux.Vars(req)
	entityId, ok := vars["id"]
	if !ok {
		http.Error(writer, "Id field must be provided!", http.StatusBadRequest)
		return err
	}

	var columnsToInsert []string
	var valuesToInsert []interface{}

	for i := 0; i < fieldsCount; i++ {
		var field = entityType.Field(i)
		if field.Name == "Id" {
			continue
		}
		columnsToInsert = append(columnsToInsert, field.Name)
		var valueToInsert interface{}
		valueToInsert = reflect.ValueOf(t).Elem().Field(i).Interface()
		if reflect.TypeOf(valueToInsert).Kind() == reflect.Ptr {
			if reflect.TypeOf(valueToInsert).Elem().Kind() == reflect.String {
				ptr := valueToInsert.(*string)
				if ptr == nil {
					valueToInsert = fmt.Sprint("null")
					continue
				}
				valueToInsert = fmt.Sprintf("'%v'", *ptr)
			}
		}

		valuesToInsert = append(valuesToInsert, valueToInsert)
	}
	var entityName = GetTypeName[T]() + "s"

	var rightPart string
	var countValuesToInsert = len(valuesToInsert)
	for i := 0; i < countValuesToInsert; i++ {
		if i > 0 {
			rightPart += ","
		}
		rightPart += columnsToInsert[i] + "=" + fmt.Sprint(valuesToInsert[i])
	}

	var query = "UPDATE %s SET %s WHERE Id = %v"
	query = fmt.Sprintf(query, entityName, rightPart, entityId)

	_, err = db.Queryx(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	return err
}

func (defaultHandler *DefaultHandler[T]) HandlePut(writer http.ResponseWriter, req *http.Request) error {
	var db = sqlx.NewDb(defaultHandler.Db, "mysql")
	var t T
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		return err
	}
	var entityType = reflect.TypeOf(t).Elem()
	var fieldsCount = entityType.NumField()

	var columnsToInsert []string
	var valuesToInsert []interface{}

	for i := 0; i < fieldsCount; i++ {
		var field = entityType.Field(i)
		if field.Name == "Id" {
			continue
		}
		columnsToInsert = append(columnsToInsert, field.Name)
		var valueToInsert interface{}
		valueToInsert = reflect.ValueOf(t).Elem().Field(i).Interface()
		if reflect.TypeOf(valueToInsert).Kind() == reflect.Ptr {
			if reflect.TypeOf(valueToInsert).Elem().Kind() == reflect.String {
				ptr := valueToInsert.(*string)
				if ptr == nil {
					valueToInsert = fmt.Sprint("null")
					continue
				}
				valueToInsert = fmt.Sprintf("'%v'", *ptr)
			}
		}

		valuesToInsert = append(valuesToInsert, valueToInsert)
	}
	var entityName = GetTypeName[T]() + "s"
	var columnsToInsertString = strings.Join(columnsToInsert, ",")
	var valuesToInsertString = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(valuesToInsert)), ","), "[]")
	var query = "INSERT INTO %s(%s) VALUES (%s);"
	query = fmt.Sprintf(query, entityName, columnsToInsertString, valuesToInsertString)

	_, err = db.Queryx(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	return err
}

func (defaultHandler *DefaultHandler[T]) HandleDelete(writer http.ResponseWriter, req *http.Request) error {
	var db = sqlx.NewDb(defaultHandler.Db, "mysql")
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		http.Error(writer, "Id field must be provided!", http.StatusBadRequest)
		return nil
	}
	var entityName = GetTypeName[T]() + "s"
	var query = "DELETE FROM %s WHERE Id = %v;"
	query = fmt.Sprintf(query, entityName, id)

	_, err := db.Queryx(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	return err
}

func GetTypeName[T any]() string {
	var entity T
	return reflect.TypeOf(entity).Elem().Name()
}
