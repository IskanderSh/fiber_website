package main

import (
	"log"

	a "website_fiber/app" // Добавление псевдонима "a" для пакета "app" из модуля сайта

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Добавление возможности работы с HTML-шаблонами
	engine := html.New("./views", ".html")

	// Запуск нового веб-сервера с заданной конфигурацией о включении шаблонизатора
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Добавление возможности отображения статических файлов (в рассматриваемом примере - CSS).
	// Пример ссылки: http://localhost:8080/index.css
	app.Static("/", "./public")

	// Работа со всеми ссылками приложения
	a.Routes(app)

	// Запуск сервера на порту :8080
	log.Fatal(app.Listen(":8080"))
}
