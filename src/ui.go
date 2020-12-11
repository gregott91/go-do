package godo

import (
	"time"

	"github.com/rivo/tview"
)

// ConfigureUI configures the tview TUI
func ConfigureUI(conn *NotesConnection) error {
	app := GenerateApplication()

	// todo handle error
	previousNotesTable, _ := populatePreviousNotesTable(GenerateTable(), conn)

	previousNotesGrid := GenerateGridWithHeader(1, 1, false)
	AddToGrid(previousNotesGrid, 0, 0, GenerateLabel("Previous Notes"))
	AddToGrid(previousNotesGrid, 1, 0, previousNotesTable)

	noteInput := GenerateInput("Add a note: ", func(input *tview.InputField) {
		note, _ := conn.CreateNote(input.GetText())
		PrependRowToTable(previousNotesTable, getRowFromNote(note)...)
		input.SetText("")
	}, func() {
		conn.CloseConnection()
		app.Stop()
	})

	parentGrid := GenerateGridWithHeader(1, 1, true)
	AddToGrid(parentGrid, 0, 0, noteInput)
	AddToGrid(parentGrid, 1, 0, previousNotesGrid)

	return SetupUI(app, parentGrid, noteInput)
}

func populatePreviousNotesTable(table *tview.Table, conn *NotesConnection) (*tview.Table, error) {
	notes, err := conn.GetNotes()

	for _, element := range notes {
		AppendRowToTable(table, getRowFromNote(element)...)
	}

	return table, err
}

func getRowFromNote(note *Note) []string {
	return []string{
		note.Timestamp.Format(time.Stamp),
		note.Text,
	}
}
