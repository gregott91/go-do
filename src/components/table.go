package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CellOptions represents the rendering options for a cell
type CellOptions struct {
	CellColor    uint64
	StartPadding int
	EndPadding   int
	ExpandCell   bool
}

// Table wraps tview.Table
type Table struct {
	Inner   *tview.Table
	Parent  *tview.Application
	Options []CellOptions
}

// CreateTable creates a table with the given options
func CreateTable(opts []CellOptions, app *Application) *Table {
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

// AddToGrid adds this table to a grid
func (table *Table) AddToGrid(grid *Grid, row int, column int) {
	grid.Inner.AddItem(table.Inner, row, column, 1, 1, 0, 0, false)
}

// AppendRow adds a row to a tview Table
func (table *Table) AppendRow(referenceValue int, rowValues ...string) {
	row := table.Inner.GetRowCount()

	table.setRowCells(row, referenceValue, rowValues...)
}

// PrependRow adds a value to a tview Table
func (table *Table) PrependRow(referenceValue int, rowValues ...string) {
	table.Inner.InsertRow(0)

	table.setRowCells(0, referenceValue, rowValues...)
}

func (table *Table) setRowCells(row int, referenceValue int, rowValues ...string) {
	for column, cell := range rowValues {
		opts := table.Options[column]

		table.setTableCell(
			referenceValue,
			row,
			column,
			getCellText(cell, opts.StartPadding, opts.EndPadding),
			opts.CellColor,
			opts.ExpandCell,
		)
	}
}

func getCellText(originalText string, startPadding int, endPadding int) string {
	paddedText := originalText

	for i := 0; i < startPadding; i++ {
		paddedText = " " + paddedText
	}

	for i := 0; i < endPadding; i++ {
		paddedText = paddedText + " "
	}

	return paddedText
}

// GetDefaultCellOptions gets the default cell options
func GetDefaultCellOptions(expand bool) CellOptions {
	return CellOptions{
		CellColor:    WrapperColorWhite,
		StartPadding: 0,
		EndPadding:   0,
		ExpandCell:   expand,
	}
}

func (table *Table) setTableCell(referenceValue int, row int, column int, text string, color uint64, expand bool) {
	cell := tview.NewTableCell(text).
		SetReference(referenceValue).
		SetTextColor(tcell.Color(color))

	if expand {
		cell = cell.SetExpansion(100)
	}

	table.Inner.SetCell(row, column, cell)
}

// GetSelectedRow gets the selected row
func (table *Table) GetSelectedRow() int {
	row, _ := table.Inner.GetSelection()
	return row
}

// GetSelectedReference gets the reference from the selected row
func (table *Table) GetSelectedReference() int {
	row := table.GetSelectedRow()
	referenceVal := table.Inner.GetCell(row, 0).GetReference()

	return referenceVal.(int)
}

// RemoveRow removes the specified row
func (table *Table) RemoveRow(row int) {
	table.Inner.RemoveRow(row)
}
