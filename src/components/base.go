package components

import "github.com/gdamore/tcell/v2"

// BaseComponent is equvalent to the tview primitive
type BaseComponent interface {
	HasFocus() bool
	SetFocus()
	AddToGrid(grid *Grid, row int, column int)
	SetRoot()
}

const (
	// WrapperColorOlive represents an olive color
	WrapperColorOlive = uint64(tcell.ColorOlive)

	// WrapperColorWhite represents a white color
	WrapperColorWhite = uint64(tcell.ColorWhite)
)
