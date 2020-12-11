package godo

import "github.com/rivo/tview"

// Application wraps tview.Application
type Application struct {
	Inner *tview.Application
}

// Table wraps tview.Table
type Table struct {
	Inner *tview.Table
}

// Grid wraps tview.Grid
type Grid struct {
	Inner *tview.Grid
}

// Label wraps tview.TextView
type Label struct {
	Inner *tview.TextView
}

// InputField wraps tview.InputField
type InputField struct {
	Inner *tview.InputField
}

// WrapTable wraps the tview table in a struct
func WrapTable(table *tview.Table) Table {
	return Table{Inner: table}
}

// WrapGrid wraps the tview grid in a struct
func WrapGrid(grid *tview.Grid) Grid {
	return Grid{Inner: grid}
}

// WrapApplication wraps the tview Application in a struct
func WrapApplication(app *tview.Application) Application {
	return Application{Inner: app}
}

// WrapLabel wraps the tview TextView in a struct
func WrapLabel(prim *tview.TextView) Label {
	return Label{Inner: prim}
}

// WrapInputField wraps the tview InputField in a struct
func WrapInputField(input *tview.InputField) InputField {
	return InputField{Inner: input}
}

// GetText returns an input's text
func (input InputField) GetText() string {
	return input.Inner.GetText()
}

// Clear clears text from an input
func (input InputField) Clear() {
	input.Inner.SetText("")
}

// Stop stops the application
func (app Application) Stop() {
	app.Inner.Stop()
}