package routes

import (
	"blogbackend/controller"
	"blogbackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/logout", controller.Logout)
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/postunik", controller.PostUnik)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/assets-image", controller.Assets)
	app.Static("/api/assets", "./assets")
}
