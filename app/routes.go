package app

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/", index) // Главная страница

	app.Get("/Student/create", createStudentGet) // Страница создания нового студента
	app.Post("/Student/create", createStudent)   // Добавление нового студента

	app.Get("/Students/:id?", readStudent) // Страница вывода данных о множестве студентов, либо одном с сохранением шаблона

	app.Get("/Student/update/:id", updateStudentGet) // Страница вывода данных для обновления данных о студенте
	app.Post("/Student/update/:id", updateStudent)   // Обновление данных о студенте

	app.Get("/Student/delete/:id", deleteStudent) // Страница, которая вызывается для удаления данных о студенте
}
