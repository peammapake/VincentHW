package app

import "github.com/peammapake/vincent-inventory-api/repository"

type App struct {
	Inventory InventoryApp
}

func NewApp(
	repo *repository.Repository,
) *App {
	return &App{
		Inventory: NewInventoryApp(repo),
	}
}