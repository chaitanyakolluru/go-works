package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserData struct {
	ID   int `json:"id"`
	User `json:"user"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []UserData

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	for _, user := range users {
		if user.User.Name == u.Name {
			return c.JSON(http.StatusConflict, "user already exists")
		}
	}

	users = append(users, UserData{ID: len(users) + 1, User: *u})
	return c.JSON(http.StatusCreated, u)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	for i, user := range users {
		if id == fmt.Sprintf("%d", user.ID) {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, id)
		}
	}

	return c.JSON(http.StatusNotFound, id)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	for i, user := range users {
		if id == fmt.Sprintf("%d", user.ID) {
			users[i].User = *u
			return c.JSON(http.StatusOK, users[i])
		}
	}

	return c.JSON(http.StatusNotFound, id)
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	for _, u := range users {
		if id == fmt.Sprintf("%d", u.ID) {
			return c.JSON(http.StatusOK, u)
		}
	}

	return c.JSON(http.StatusNotFound, id)
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "status: ok")
}

func main() {
	e := echo.New()

	e.GET("/healthz", healthCheck)
	e.POST("/users", saveUser)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.DELETE("/users/:id", deleteUser)
	e.PUT("/users/:id", updateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
