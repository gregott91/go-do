package components

import "github.com/rivo/tview"

// LabelOptions is the configuration options for the Label
type LabelOptions struct {
	Text   string
	Center bool
}

// Label wraps tview.TextView
type Label struct {
	Inner  *tview.TextView
	Parent *tview.Application
}

// CreateLabel creates a label with the given text
func CreateLabel(opts LabelOptions, app *Application) *Label {
	label := &Label{
		Inner: tview.NewTextView().
			SetText(opts.Text),
		Parent: app.Inner,
	}

	if opts.Center {
		label.Inner.SetTextAlign(tview.AlignCenter)
	} else {
		label.Inner.SetTextAlign(tview.AlignLeft)
	}

	return label
}

// AddToGrid adds this label to a grid
func (component *Label) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(component.Inner, row, column, 1, 1, 0, 0, false)
}
