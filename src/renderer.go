package godo

import (
	"go-do/src/components"
)

// SetupUI confiugures the root and focus of the TUI
func SetupUI(app components.Application, root components.Grid, focus components.InputField) error {
	return app.Inner.SetRoot(root.Inner, true).SetFocus(focus.Inner).Run()
}
