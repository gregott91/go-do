package components

import "github.com/rivo/tview"

// Label wraps tview.TextView
type Label struct {
	Inner  *tview.TextView
	Parent *tview.Application
}

// CreateLabel creates a label with the given text
func CreateLabel(text string, app *Application) *Label {
	return &Label{
		Inner: tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text),
		Parent: app.Inner,
	}
}

// AddToGrid adds this label to a grid
func (component *Label) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(component.Inner, row, column, 1, 1, 0, 0, false)
}
