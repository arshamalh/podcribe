package api

import "github.com/labstack/echo/v4"

type api struct {
	router *echo.Echo
	// Add db here
}

func New() *api {
	router := echo.New()
	return &api{
		router: router,
	}
}

func (a *api) StartBlocking() {
	a.router.Start("localhost:80")
}

// Not implemented
func (a *api) StartAsync() {}
