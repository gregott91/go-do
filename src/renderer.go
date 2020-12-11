package godo

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// GenerateApplication generates the tview Application
func GenerateApplication() *tview.Application {
	return tview.NewApplication()
}

// GenerateGridWithHeader generates the tview Grid with a single header row
func GenerateGridWithHeader(rowCount int, columnCount int, border bool) *tview.Grid {
	rows := []int{1}
	for i := 1; i < rowCount; i++ {
		rows = append(rows, i)
	}

	columns := []int{}
	for i := 1; i < columnCount; i++ {
		columns = append(columns, i)
	}

	return tview.NewGrid().
		SetRows(rows...).
		SetColumns(columns...).
		SetBorders(border)
}

// GenerateLabel generates a tview Primitive from a string
func GenerateLabel(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

// AddToGrid adds a value to a grid
func AddToGrid(grid *tview.Grid, row int, column int, item tview.Primitive) {
	grid.AddItem(item, row, column, 1, 1, 0, 0, false)
}

// SetupUI confiugures the root and focus of the TUI
func SetupUI(app *tview.Application, root tview.Primitive, focus tview.Primitive) error {
	return app.SetRoot(root, true).SetFocus(focus).Run()
}

// GenerateInput generates an input field
func GenerateInput(label string, enterFunc func(input *tview.InputField), closeFunc func()) *tview.InputField {
	input := tview.NewInputField().
		SetLabel(label).
		SetFieldBackgroundColor(tcell.ColorBlack)

	handleInputDone := func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			enterFunc(input)
		case tcell.KeyEscape:
			closeFunc()
		}
	}

	return input.SetDoneFunc(handleInputDone)
}

// GenerateTable generates a tview Table
func GenerateTable() *tview.Table {
	table := tview.NewTable().
		SetBorders(false)

	table.
		Select(0, 0).
		SetFixed(1, 1).
		SetSelectedFunc(func(row int, column int) {
			table.GetCell(row, column).SetTextColor(tcell.ColorRed)
			table.SetSelectable(false, false)
		})

	return table
}

// AppendRowToTable adds a row to a tview Table
func AppendRowToTable(table *tview.Table, rowValues ...string) {
	row := table.GetRowCount()

	setRowCells(table, row, rowValues...)
}

// PrependRowToTable adds a value to a tview Table
func PrependRowToTable(table *tview.Table, rowValues ...string) {
	table.InsertRow(0)

	setRowCells(table, 0, rowValues...)
}

func setRowCells(table *tview.Table, row int, rowValues ...string) {
	for column, cell := range rowValues {
		setTableCell(table, row, column, cell)
	}
}

func setTableCell(table *tview.Table, row int, column int, text string) {
	table.SetCell(row, column,
		tview.NewTableCell(text).
			SetTextColor(tcell.ColorWhite))
}
