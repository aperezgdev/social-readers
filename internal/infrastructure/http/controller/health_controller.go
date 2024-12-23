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
	_, err := w.Write([]byte("OK"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
