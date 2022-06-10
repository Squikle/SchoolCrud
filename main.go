package main

import (
	"DbProj/DAL"
	"DbProj/handlers"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3307)/DbProj")

	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/school/{id}", handlers.NewDefaultHandler[*DAL.School](db).Handle)
	r.HandleFunc("/school", handlers.NewDefaultHandler[*DAL.School](db).Handle)

	r.HandleFunc("/car/{id}", handlers.NewDefaultHandler[*DAL.Car](db).Handle)
	r.HandleFunc("/car", handlers.NewDefaultHandler[*DAL.Car](db).Handle)

	r.HandleFunc("/group/{id}", handlers.NewDefaultHandler[*DAL.Group](db).Handle)
	r.HandleFunc("/group", handlers.NewDefaultHandler[*DAL.Group](db).Handle)

	r.HandleFunc("/instructor/{id}", handlers.NewDefaultHandler[*DAL.Instructor](db).Handle)
	r.HandleFunc("/instructor", handlers.NewDefaultHandler[*DAL.Instructor](db).Handle)

	r.HandleFunc("/lesson/{id}", handlers.NewDefaultHandler[*DAL.Lessons](db).Handle)
	r.HandleFunc("/lesson", handlers.NewDefaultHandler[*DAL.Lessons](db).Handle)

	r.HandleFunc("/payment/{id}", handlers.NewDefaultHandler[*DAL.Payments](db).Handle)
	r.HandleFunc("/payment", handlers.NewDefaultHandler[*DAL.Payments](db).Handle)

	r.HandleFunc("/person/{id}", handlers.NewDefaultHandler[*DAL.Person](db).Handle)
	r.HandleFunc("/person", handlers.NewDefaultHandler[*DAL.Person](db).Handle)

	r.HandleFunc("/practiseLesson/{id}", handlers.NewDefaultHandler[*DAL.PracticeLesson](db).Handle)
	r.HandleFunc("/practiseLesson", handlers.NewDefaultHandler[*DAL.PracticeLesson](db).Handle)

	r.HandleFunc("/student/{id}", handlers.NewDefaultHandler[*DAL.Student](db).Handle)
	r.HandleFunc("/student", handlers.NewDefaultHandler[*DAL.Student](db).Handle)

	r.HandleFunc("/studentGroup/{id}", handlers.NewDefaultHandler[*DAL.StudentsGroup](db).Handle)
	r.HandleFunc("/studentGroup", handlers.NewDefaultHandler[*DAL.StudentsGroup](db).Handle)

	r.HandleFunc("/theoryLesson/{id}", handlers.NewDefaultHandler[*DAL.TheoryLesson](db).Handle)
	r.HandleFunc("/theoryLesson", handlers.NewDefaultHandler[*DAL.TheoryLesson](db).Handle)

	r.HandleFunc("/topic/{id}", handlers.NewDefaultHandler[*DAL.Topic](db).Handle)
	r.HandleFunc("/topic", handlers.NewDefaultHandler[*DAL.Topic](db).Handle)

	err = http.ListenAndServe(":5462", r)

	if err != nil {
		panic(err.Error())
	}
}
