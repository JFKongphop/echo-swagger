// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /

package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/JFKongphop/echo-swagger/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// HTTPError represents an error that occurred while handling a request.
// @Description HTTP error
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// getUser is an example handler that demonstrates the use of HTTPError.
// @Summary Show a user
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/{id} [get]
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "User "+id)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users/:id", getUser)

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
