package godo

import (
	"time"
)

// ConfigureUI configures the TUI
func ConfigureUI(conn *NotesConnection) error {
	app := GenerateApplication()

	// todo handle error
	previousNotesTable, _ := GenerateTable(WrapperColorOlive).
		populatePreviousNotesTable(conn)

	previousNotesGrid := GenerateGridWithHeader(1, 1, false)
	label := GenerateLabel("Previous Notes")
	AddLabelToGrid(previousNotesGrid, 0, 0, &label)
	AddTableToGrid(previousNotesGrid, 1, 0, &previousNotesTable)

	noteInput := GenerateInput("Add a note: ", func(input InputField) {
		note, _ := conn.CreateNote(input.GetText())
		PrependRowToTable(previousNotesTable, getRowFromNote(note)...)
		input.Clear()
	}, func() {
		conn.CloseConnection()
		app.Stop()
	})

	parentGrid := GenerateGridWithHeader(1, 1, true)
	AddInputToGrid(parentGrid, 0, 0, &noteInput)
	AddGridToGrid(parentGrid, 1, 0, &previousNotesGrid)

	return SetupUI(app, parentGrid, noteInput)
}

func (table Table) populatePreviousNotesTable(conn *NotesConnection) (Table, error) {
	notes, err := conn.GetNotes()

	for _, element := range notes {
		AppendRowToTable(table, getRowFromNote(element)...)
	}

	return table, err
}

func getRowFromNote(note *Note) []string {
	return []string{
		note.Timestamp.Format(time.Stamp) + "  ",
		note.Text,
	}
}
