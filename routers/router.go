package routers

import (
    "todo_backend/handlers"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
    router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
    return router
}
