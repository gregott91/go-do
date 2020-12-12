package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Application wraps tview.Application
type Application struct {
	Inner *tview.Application
}

// CreateApplication generates the tview Application
func CreateApplication() *Application {
	return &Application{Inner: tview.NewApplication()}
}

// Stop stops the application
func (app *Application) Stop() {
	app.Inner.Stop()
}

// ConfigureAppShortcuts configues the shortcuts for the application
func (app *Application) ConfigureAppShortcuts(handler func(uint64)) {
	app.Inner.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		handler(uint64(event.Key()))
		return event
	})
}

// Run runs the application
func (app *Application) Run() error {
	return app.Inner.Run()
}
