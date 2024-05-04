// Package web contains a small web framework extension.
package web

import (
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux"
)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	/*mux  */ *httptreemux.ContextMux
	//otmux    http.Handler
	shutdown chan os.Signal
	//mw       []Middleware
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	//
	mux := httptreemux.NewContextMux()

	return &App{
		ContextMux: mux,
		shutdown:   shutdown,
	}
}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
