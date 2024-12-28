package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var healthController = NewHealthController()

func TestHelthController(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", &bytes.Buffer{})
	w := httptest.NewRecorder()
	healthController.GetHealth(w, r)

	var response = w.Body.String()

	assert.Equal(t, "OK", response)
	assert.Equal(t, http.StatusOK, w.Code)
}