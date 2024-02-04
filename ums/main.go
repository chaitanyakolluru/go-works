package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type UserData struct {
	gorm.Model
	User `json:"user"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []UserData

type Context struct {
	echo.Context
	DB *gorm.DB
}

func saveUser(cc Context) error {
	u := new(User)
	if err := cc.Bind(u); err != nil {
		return err
	}

	for _, user := range users {
		if user.User.Name == u.Name {
			return cc.JSON(http.StatusConflict, "user already exists")
		}
	}

	cc.DB.Create(&UserData{User: *u})

	return cc.JSON(http.StatusCreated, u)
}

func deleteUser(cc Context) error {
	id := cc.Param("id")

	for i, user := range users {
		if id == fmt.Sprintf("%d", user.ID) {
			users = append(users[:i], users[i+1:]...)
			return cc.JSON(http.StatusOK, id)
		}
	}

	return cc.JSON(http.StatusNotFound, "user deleted")
}

func updateUser(cc Context) error {
	id := cc.Param("id")
	u := new(User)
	if err := cc.Bind(u); err != nil {
		return err
	}

	for i, user := range users {
		if id == fmt.Sprintf("%d", user.ID) {
			users[i].User = *u
			return cc.JSON(http.StatusOK, users[i])
		}
	}

	return cc.JSON(http.StatusNotFound, "user not found")
}

func getUsers(cc Context) error {
	return cc.JSON(http.StatusOK, users)
}

func getUser(cc Context) error {
	id := cc.Param("id")

	for _, u := range users {
		if id == fmt.Sprintf("%d", u.ID) {
			return cc.JSON(http.StatusOK, u)
		}
	}

	return cc.JSON(http.StatusNotFound, "user not found")
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "status: ok")
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&UserData{})
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c, db}
			return next(cc)
		}
	})

	e.GET("/healthz", healthCheck)
	e.POST("/users", saveUser)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.DELETE("/users/:id", deleteUser)
	e.PUT("/users/:id", updateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
