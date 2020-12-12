package components

import "github.com/rivo/tview"

// Grid wraps tview.Grid
type Grid struct {
	Inner  *tview.Grid
	Parent *tview.Application
}

type GridOptions struct {
	Rows         []int
	Columns      []int
	HasBorders   bool
	HasHeaderRow bool
}

func GetDefaultGridOptions(hasBorder bool) GridOptions {
	return GridOptions{
		Rows:         []int{0},
		Columns:      []int{0},
		HasBorders:   hasBorder,
		HasHeaderRow: true,
	}
}

func CreateGrid(opts GridOptions, app *Application) *Grid {
	var finalRows []int

	if opts.HasHeaderRow {
		finalRows = []int{1}
	}

	finalRows = append(finalRows, opts.Rows...)

	return &Grid{
		Inner: tview.NewGrid().
			SetRows(finalRows...).
			SetColumns(opts.Columns...).
			SetBorders(opts.HasBorders),
		Parent: app.Inner,
	}
}

func (component *Grid) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(component.Inner, row, column, 1, 1, 0, 0, false)
}

func (component *Grid) SetRoot() {
	component.Parent.SetRoot(component.Inner, true)
}
