package godo

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// GenerateApplication generates the tview Application
func GenerateApplication() Application {
	return WrapApplication(tview.NewApplication())
}

// GenerateGridWithHeader generates the tview Grid with a single header row
func GenerateGridWithHeader(rowCount int, columnCount int, border bool) Grid {
	rows := []int{1}
	for i := 1; i < rowCount; i++ {
		rows = append(rows, i)
	}

	columns := []int{}
	for i := 1; i < columnCount; i++ {
		columns = append(columns, i)
	}

	return WrapGrid(
		tview.NewGrid().
			SetRows(rows...).
			SetColumns(columns...).
			SetBorders(border))
}

// GenerateLabel generates a tview Primitive from a string
func GenerateLabel(text string) Label {
	return WrapLabel(
		tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text))
}

// AddLabelToGrid adds a value to a grid
func AddLabelToGrid(grid Grid, row int, column int, item *Label) {
	grid.Inner.AddItem(item.Inner, row, column, 1, 1, 0, 0, false)
}

// AddInputToGrid adds a value to a grid
func AddInputToGrid(grid Grid, row int, column int, item *InputField) {
	grid.Inner.AddItem(item.Inner, row, column, 1, 1, 0, 0, false)
}

// AddTableToGrid adds a value to a grid
func AddTableToGrid(grid Grid, row int, column int, item *Table) {
	grid.Inner.AddItem(item.Inner, row, column, 1, 1, 0, 0, false)
}

// AddGridToGrid adds a value to a grid
func AddGridToGrid(grid Grid, row int, column int, item *Grid) {
	grid.Inner.AddItem(item.Inner, row, column, 1, 1, 0, 0, false)
}

// SetupUI confiugures the root and focus of the TUI
func SetupUI(app Application, root Grid, focus InputField) error {
	return app.Inner.SetRoot(root.Inner, true).SetFocus(focus.Inner).Run()
}

// GenerateInput generates an input field
func GenerateInput(label string, enterFunc func(input InputField), closeFunc func()) InputField {
	input := tview.NewInputField().
		SetLabel(label).
		SetFieldBackgroundColor(tcell.ColorBlack)

	handleInputDone := func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			enterFunc(WrapInputField(input))
		case tcell.KeyEscape:
			closeFunc()
		}
	}

	return WrapInputField(input.SetDoneFunc(handleInputDone))
}

// GenerateTable generates a tview Table
func GenerateTable(firstCellColor uint64) Table {
	table := tview.NewTable().
		SetBorders(false)

	table.
		Select(0, 0).
		SetFixed(1, 1).
		SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		})

	return WrapTable(table, firstCellColor)
}

// AppendRowToTable adds a row to a tview Table
func AppendRowToTable(table Table, rowValues ...string) {
	row := table.Inner.GetRowCount()

	setRowCells(table, row, rowValues...)
}

// PrependRowToTable adds a value to a tview Table
func PrependRowToTable(table Table, rowValues ...string) {
	table.Inner.InsertRow(0)

	setRowCells(table, 0, rowValues...)
}

func setRowCells(table Table, row int, rowValues ...string) {
	for column, cell := range rowValues {
		color := uint64(tcell.ColorWhite)
		if column == 0 {
			color = table.Options.FirstCellColor
		}

		setTableCell(table, row, column, cell, color)
	}
}

func setTableCell(table Table, row int, column int, text string, color uint64) {
	table.Inner.SetCell(row, column,
		tview.NewTableCell(text).
			SetTextColor(tcell.Color(color)))
}
