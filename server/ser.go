package main

import (
    "fmt"
    "log"
    "net/http"
)

// Обработчик запросов
func handler(w http.ResponseWriter, r *http.Request) {
    // w - для отправки ответа клиенту
    // r - данные запроса
    
    fmt.Fprintf(w, "Привет! Ты запросил: %s", r.URL.Path)
}

func main() {
    // Правильно: HandleFunc (а не Handeler)
    http.HandleFunc("/", handler)
    
    fmt.Println("Сервер запущен на http://localhost:8000")
    
    // Запуск сервера на порту 8080
    log.Fatal(http.ListenAndServe(":8000", nil))
}