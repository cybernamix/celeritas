package middleware

import (
	"test5/data"

	"github.com/cybernamix/celeritas"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}