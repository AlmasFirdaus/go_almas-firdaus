package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var resUser User

	for _, user := range users {
		if user.Id == id {
			resUser = user
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"users":    resUser,
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var newUser []User

	for _, user := range users {
		if user.Id != id {
			newUser = append(newUser, user)
		}
	}
	users = newUser
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"users":    users,
	})
}

// // update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, user := range users {
		if user.Id == id {
			userUpdate := user
			userUpdate.Name = "name has changed hehehe"
			users[i] = userUpdate
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success updated user",
		"users":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{Name: "almas", Email: "almas@gmail.com", Password: "asd123"}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
