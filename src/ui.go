package godo

import (
	"go-do/src/components"
	"sort"
	"time"
)

// ConfigureUI configures the TUI
func ConfigureUI(conn *NotesConnection) error {
	app := components.CreateApplication()

	// todo handle error
	previousNotesTable, _ := getPopulatedTable(app, conn)

	previousNotesGrid := components.CreateGrid(components.GetDefaultGridOptions(false), app)
	label := components.CreateLabel("Previous Notes", app)
	label.AddToGrid(previousNotesGrid, 0, 0)
	previousNotesTable.AddToGrid(previousNotesGrid, 1, 0)

	noteInput := components.CreateInputField(components.InputOptions{
		Label:        "Add a note:",
		LabelPadding: 1,
		EnterFunc: func(input components.InputField) {
			note, _ := conn.CreateNote(input.GetText())
			previousNotesTable.PrependRow(getRowFromNote(note)...)
			input.Clear()
		},
		CloseFunc: func() {
			conn.CloseConnection()
			app.Stop()
		},
	}, app)

	parentGrid := components.CreateGrid(components.GetDefaultGridOptions(true), app)
	noteInput.AddToGrid(parentGrid, 0, 0)
	previousNotesGrid.AddToGrid(parentGrid, 1, 0)

	configureUIShortcuts(app, noteInput, previousNotesTable)

	parentGrid.SetRoot()
	noteInput.SetFocus()

	return app.Run()
}

func configureUIShortcuts(app *components.Application, input *components.InputField, notesTable *components.Table) {
	app.ConfigureAppShortcuts(func(keyCode uint64) {
		if keyCode == components.KeyCtrlS {
			if input.HasFocus() {
				notesTable.EnableSelection()
				notesTable.SetFocus()
			} else if notesTable.HasFocus() {
				notesTable.DisableSelection()
				input.SetFocus()
			}
		}
	})
}

func getPopulatedTable(app *components.Application, conn *NotesConnection) (*components.Table, error) {
	table := components.CreateTable([]components.CellOptions{
		{
			CellColor:    components.WrapperColorOlive,
			StartPadding: 1,
			EndPadding:   2,
		},
		components.GetDefaultCellOptions(),
	}, app)

	notes, err := conn.GetNotes()
	sort.Sort(byAge(notes))
	for _, element := range notes {
		table.AppendRow(getRowFromNote(element)...)
	}

	return table, err
}

func getRowFromNote(note *Note) []string {
	return []string{
		note.Timestamp.Format(time.Stamp),
		note.Text,
	}
}
