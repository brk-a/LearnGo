package routes

import (
	"log"
	"net/http"
	"fmt"
	"strconv"

	"html/template"
	"htmx/model"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter)  {
	todos, err := model.GetAllTodos()

	if err!=nil {
		fmt.Println("Could not get all to-do items from db", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err!=nil {
		fmt.Println("Could not execute template", err)
	}
}

func index(w http.ResponseWriter, r *http.Request)  {
	// sendTodos(w)
	todos, err := model.GetAllTodos()

	if err!=nil {
		fmt.Println("Could not get all to-do items from db", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err!=nil {
		fmt.Println("Could not execute template", err)
	}
}
func markTodo(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err!=nil {
		fmt.Println("Could not parse id", err)
	}

	err = model.MarkTodo(id)
	if(err!=nil) {
		fmt.Println("Could not update to-do", err)
	}

	sendTodos(w)
}

func createTodo(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err!=nil {
		fmt.Println("Error parsing form", err)
	}

	err = model.CreateTodo(r.FormValue("todo"))
	if err!=nil {
		fmt.Println("Could not crete to-do", err)
	}

	sendTodos(w)
}
func deleteTodo(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err!=nil {
		fmt.Println("Could not parse id", err)
	}

	err = model.DeleteTodo(id)
	if err!=nil {
		fmt.Println("Could not delete", err)
	}

	sendTodos(w)
}

SetupAndRun()  {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000",  mux))

}
