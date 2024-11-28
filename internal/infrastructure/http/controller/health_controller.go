package controller

import "net/http"

type HealthController struct {
	pattern string
}

func NewHealthController() HealthController {
	return HealthController{
		"health",
	}
}

func (HealthController) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
