// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/kamogelosekhukhune777/bookings/foundation/web"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs an http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {

	// Construct the web.App which holds all routes.
	app := web.NewApp(
		cfg.Shutdown,
	)
	app.Handle(http.MethodGet, "/test", Test)

	return app
}

// Test handler is for development.
func Test(w http.ResponseWriter, r *http.Request) {

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(&status)
}
