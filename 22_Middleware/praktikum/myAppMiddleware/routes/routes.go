package routes

import (
	"myAppMiddleware/constants"
	"myAppMiddleware/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	// users
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	//jwt group
	r := e.Group("/jwt")
	r.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserDetailControllers)

	// books
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
