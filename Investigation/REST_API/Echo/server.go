package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/* e.GET("/users/:id", getUser) */
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// teamとmemberというクエリ文字列を取得する
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func main() {
	e := echo.New() // Echoのインスタンスの作成

	/* Routing */
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)

	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	/* GETリクエスト,  */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
