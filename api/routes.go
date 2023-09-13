package api

import (
	"podcribe/api/handlers"
)

func (a *api) RegisterRoutes() {
	// A route for uploading the file
	// First, check for extension (front also check it too)
	// If it was valid, depending .mp3 or wav, make a new manager and follow a path

	api := a.router.Group("api")
	// A route for doing the flow, user decide whether to download or ... using some buttons
	api.POST("process", handlers.Process())
}
