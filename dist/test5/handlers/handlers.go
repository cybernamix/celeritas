package handlers

import (
	"test5/data"
	"net/http"
	"time"

	"github.com/cybernamix/celeritas"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

//Can use the convenience.go helpers of h.render rather than the full h.App.Render.Page...

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
