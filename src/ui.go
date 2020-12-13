package godo

import (
	"go-do/src/components"
	"sort"
	"time"
)

type shortcutConfig struct {
	config     *Config
	observable *ShortcutObservable
}

// ConfigureUI configures the TUI
func ConfigureUI(conn *NotesConnection, config *Config) (*components.Application, error) {
	shortcuts := &shortcutConfig{config: config, observable: &ShortcutObservable{Handlers: make(map[uint64]func())}}

	app := createApplication(shortcuts, conn)

	previousNotesTable, err := getPopulatedTable(app, conn, shortcuts)

	if err != nil {
		return &components.Application{}, err
	}

	previousNotesGrid := getNoteGrid(app, previousNotesTable)
	noteInput := getNoteInput(conn, previousNotesTable, app, shortcuts)

	parentGrid := components.CreateGrid(components.GetDefaultGridOptions(true), app)
	noteInput.AddToGrid(parentGrid, 0, 0)
	previousNotesGrid.AddToGrid(parentGrid, 1, 0)

	parentGrid.SetRoot()
	noteInput.SetFocus()

	return app, app.Run()
}

func createApplication(shortcuts *shortcutConfig, conn *NotesConnection) *components.Application {
	app := components.CreateApplication()
	app.ConfigureInputCapture(shortcuts.observable.Trigger)

	shortcuts.observable.Register(shortcuts.config.Shortcuts.Close, func() {
		conn.CloseConnection()
		app.Stop()
	})

	return app
}

func getNoteGrid(app *components.Application, previousNotesTable *components.Table) *components.Grid {
	previousNotesGrid := components.CreateGrid(components.GetDefaultGridOptions(false), app)
	label := components.CreateLabel("Previous Notes", app)
	label.AddToGrid(previousNotesGrid, 0, 0)
	previousNotesTable.AddToGrid(previousNotesGrid, 1, 0)

	return previousNotesGrid
}

func getNoteInput(conn *NotesConnection, notesTable *components.Table, app *components.Application, shortcuts *shortcutConfig) *components.InputField {
	field := components.CreateInputField(components.InputOptions{
		Label:        "Add a note:",
		LabelPadding: 1,
		EnterFunc: func(input components.InputField) {
			note, _ := conn.CreateNote(input.GetText())
			notesTable.PrependRow(note.ID, getRowFromNote(note)...)
			input.Clear()
		},
	}, app)

	shortcuts.observable.Register(shortcuts.config.Shortcuts.Switch, func() {
		if field.HasFocus() {
			notesTable.EnableSelection()
			notesTable.SetFocus()
		} else if notesTable.HasFocus() {
			notesTable.DisableSelection()
			field.SetFocus()
		}
	})

	return field
}

func getPopulatedTable(app *components.Application, conn *NotesConnection, shortcuts *shortcutConfig) (*components.Table, error) {
	table := components.CreateTable([]components.CellOptions{
		{
			CellColor:    components.WrapperColorOlive,
			StartPadding: 1,
			EndPadding:   2,
			ExpandCell:   false,
		},
		components.GetDefaultCellOptions(true),
	}, app)

	notes, err := conn.GetNotes()
	sort.Sort(byAge(notes))
	for _, element := range notes {
		table.AppendRow(element.ID, getRowFromNote(element)...)
	}

	shortcuts.observable.Register(shortcuts.config.Shortcuts.Delete, func() {
		if table.HasFocus() {
			id := table.GetSelectedReference()
			conn.RemoveNote(id)
			table.RemoveRow(table.GetSelectedRow())
		}
	})

	return table, err
}

func getRowFromNote(note *Note) []string {
	return []string{
		note.Timestamp.Format(time.Stamp),
		note.Text,
	}
}
