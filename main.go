package main

import (
    "log"
    "net/http"
    "todo_backend/database"
    "todo_backend/routers"
)

func enableCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    database.Connect()
    router := routers.SetupRouter()
	// router.Use(enableCors)
	corsEnabledRouter := enableCors(router)
	log.Println("Udah starting server on :8080")
    // log.Fatal(http.ListenAndServe(":8080", router))
	log.Fatal(http.ListenAndServe(":8080", corsEnabledRouter))
}
