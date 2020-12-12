package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// TableOptions represents options for a tview.Table
type TableOptions struct {
	FirstCellColor uint64
}

// Table wraps tview.Table
type Table struct {
	Inner   *tview.Table
	Options TableOptions
	Parent  *tview.Application
}

func CreateTable(opts TableOptions, app *Application) *Table {
	return &Table{
		Inner: tview.NewTable().
			SetBorders(false).
			SetSelectable(false, false),
		Options: opts,
		Parent:  app.Inner,
	}
}

// HasFocus returns true if the table has keyboard focus
func (table *Table) HasFocus() bool {
	return table.Inner.HasFocus()
}

// SetFocus sets the focus
func (table *Table) SetFocus() {
	table.Parent.SetFocus(table.Inner)
}

// EnableSelection enables selectability
func (table *Table) EnableSelection() {
	table.Inner.SetSelectable(true, false)
}

// DisableSelection disables selectability
func (table *Table) DisableSelection() {
	table.Inner.SetSelectable(false, false)
}

func (component *Table) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(component.Inner, row, column, 1, 1, 0, 0, false)
}

// AppendRow adds a row to a tview Table
func (table *Table) AppendRow(rowValues ...string) {
	row := table.Inner.GetRowCount()

	table.setRowCells(row, rowValues...)
}

// PrependRow adds a value to a tview Table
func (table *Table) PrependRow(rowValues ...string) {
	table.Inner.InsertRow(0)

	table.setRowCells(0, rowValues...)
}

func (table *Table) setRowCells(row int, rowValues ...string) {
	for column, cell := range rowValues {
		color := uint64(tcell.ColorWhite)
		if column == 0 {
			color = table.Options.FirstCellColor
		}

		expand := false
		if column == (len(rowValues) - 1) {
			expand = true
		}

		table.setTableCell(row, column, cell, color, expand)
	}
}

func (table *Table) setTableCell(row int, column int, text string, color uint64, expand bool) {
	cell := tview.NewTableCell(text).
		SetTextColor(tcell.Color(color))

	if expand {
		cell = cell.SetExpansion(0)
	}

	table.Inner.SetCell(row, column, cell)
}
