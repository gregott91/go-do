package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InputOptions struct {
	Label        string
	LabelPadding int
	EnterFunc    func(input InputField)
	CloseFunc    func()
}

// InputField wraps tview.InputField
type InputField struct {
	Inner  *tview.InputField
	Parent *tview.Application
}

// CreateInputField generates an input field
func CreateInputField(opts InputOptions, app *Application) *InputField {
	labelText := opts.Label

	for i := 0; i < opts.LabelPadding; i++ {
		labelText = " " + labelText + " "
	}

	input := tview.NewInputField().
		SetLabel(labelText).
		SetFieldBackgroundColor(tcell.ColorBlack)

	wrapped := InputField{
		Inner:  input,
		Parent: app.Inner,
	}

	handleInputDone := func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			opts.EnterFunc(wrapped)
		case tcell.KeyEscape:
			opts.CloseFunc()
		}
	}

	wrapped.Inner.SetDoneFunc(handleInputDone)

	return &wrapped
}

// GetText returns an input's text
func (input *InputField) GetText() string {
	return input.Inner.GetText()
}

// Clear clears text from an input
func (input *InputField) Clear() {
	input.Inner.SetText("")
}

// HasFocus returns true if the input has keyboard focus
func (input *InputField) HasFocus() bool {
	return input.Inner.HasFocus()
}

// SetFocus sets the focus
func (input *InputField) SetFocus() {
	input.Parent.SetFocus(input.Inner)
}

func (component *InputField) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(component.Inner, row, column, 1, 1, 0, 0, false)
}
