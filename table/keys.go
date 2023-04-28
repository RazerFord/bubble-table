package table

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the keybindings for the table when it's focused.
type KeyMap struct {
	RowDown key.Binding
	RowUp   key.Binding

	RowSelectToggle key.Binding

	PageDown  key.Binding
	PageUp    key.Binding
	PageFirst key.Binding
	PageLast  key.Binding

	Help key.Binding

	Install key.Binding
	Delete  key.Binding
	Update  key.Binding

	// Filter allows the user to start typing and filter the rows.
	Filter key.Binding

	// FilterBlur is the key that stops the user's input from typing into the filter.
	FilterBlur key.Binding

	// FilterClear will clear the filter while it's blurred.
	FilterClear key.Binding

	// ScrollRight will move one column to the right when overflow occurs.
	ScrollRight key.Binding

	// ScrollLeft will move one column to the left when overflow occurs.
	ScrollLeft key.Binding
}

// DefaultKeyMap returns a set of sensible defaults for controlling a focused table.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		RowDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		RowUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		RowSelectToggle: key.NewBinding(
			key.WithKeys(" ", "enter"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("right", "l", "pgdown"),
			key.WithHelp("→/l", "next page"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("left", "h", "pgup"),
			key.WithHelp("←/h", "previous page"),
		),
		PageFirst: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g", "first page"),
		),
		PageLast: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G", "last page"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "Help"),
		),
		Install: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "install"),
		),
		Update: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "update"),
		),
		Delete: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete"),
		),
		Filter: key.NewBinding(
			key.WithKeys("/"),
		),
		FilterBlur: key.NewBinding(
			key.WithKeys("enter", "esc"),
		),
		FilterClear: key.NewBinding(
			key.WithKeys("esc"),
		),
		ScrollRight: key.NewBinding(
			key.WithKeys("shift+right"),
		),
		ScrollLeft: key.NewBinding(
			key.WithKeys("shift+left"),
		),
	}
}
