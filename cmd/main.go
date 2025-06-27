package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "go-task-api/internal/task"
)

func main() {
    r := mux.NewRouter()

    // Registra os handlers da feature "task"
    task.RegisterHandlers(r)

    log.Println("🚀 Servidor iniciado em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
