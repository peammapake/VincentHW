package server

import "github.com/peammapake/vincent-inventory-api/internal/app"

type Server struct {
	App *app.App
}

func NewServer(app *app.App) *Server {
	return &Server{
		App: app,
	}
}
