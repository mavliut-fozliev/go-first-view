package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, мир!")
}

func main() {
	http.HandleFunc("/", handler)     // Регистрируем обработчик для корня "/"
	http.ListenAndServe(":8081", nil) // Запускаем сервер на порту 8081
}
