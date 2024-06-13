package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/otavio-Pucharelli/filhos-da-luz/internal/config"
)

var app *config.AppConfig

// NewHelpers sets the config for the helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError is a helper function that sends a client error response
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError is a helper function that sends a server error response
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println("Server error:", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
