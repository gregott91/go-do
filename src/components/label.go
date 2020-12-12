package components

import "github.com/rivo/tview"

// Label wraps tview.TextView
type Label struct {
	Inner  *tview.TextView
	Parent *tview.Application
}

func CreateLabel(text string, app *Application) *Label {
	return &Label{
		Inner: tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text),
		Parent: app.Inner,
	}
}

func (component *Label) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(component.Inner, row, column, 1, 1, 0, 0, false)
}
