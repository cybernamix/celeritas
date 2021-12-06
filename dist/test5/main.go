package main

import (
	"test5/data"
	"test5/handlers"
	"test5/middleware"

	"github.com/cybernamix/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {

	c := initApplication()
	c.App.ListenAndServe()

}
