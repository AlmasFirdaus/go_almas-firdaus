package routes

import (
	"myAppTesting/constants"
	"myAppTesting/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	// login
	e.POST("/login", controllers.LoginUsersController)

	// users
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// books
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	//jwt group
	r := e.Group("/jwt")
	r.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserDetailControllers)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.GET("/books", controllers.GetBooksController)
	r.GET("/books/:id", controllers.GetBookController)
	r.POST("/books", controllers.CreateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
