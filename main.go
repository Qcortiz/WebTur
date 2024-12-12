package main

import (
	"log"
	"net/http"
)

func main() {
	// Обработка статических файлов (например, изображения)

	// Обработчик для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Загружаем файл главной страницы
		http.ServeFile(w, r, "index.html")
	})

	// Обработчик для страницы "О нас"
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// Загружаем файл "onas.html"
		http.ServeFile(w, r, "onas.html")
	})

	// Обработчик для страницы "Купить билет"
	http.HandleFunc("/buy-ticket", func(w http.ResponseWriter, r *http.Request) {
		// Загружаем файл "buy-ticket.html"
		http.ServeFile(w, r, "buy-ticket.html")
	})

	// Обработчик для страницы "Популярные поездки"
	http.HandleFunc("/popular-trips", func(w http.ResponseWriter, r *http.Request) {
		// Загружаем файл "popular-trips.html"
		http.ServeFile(w, r, "popular-trips.html")
	})

	// Обработчик для страницы "Техническая поддержка"
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		// Загружаем файл "contact.html"
		http.ServeFile(w, r, "contact.html")
	})

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
