package table

import "github.com/charmbracelet/lipgloss"

// HeaderStyle sets the style to apply to the header text, such as color or bold.
func (m Model) HeaderStyle(style lipgloss.Style) Model {
	m.headerStyle = style.Copy()

	return m
}

// WithRows sets the rows to show as data in the table.
func (m Model) WithRows(rows []Row) Model {
	m.rows = rows

	return m
}

// WithKeyMap sets the key map to use for controls when focused.
func (m Model) WithKeyMap(keyMap KeyMap) Model {
	m.keyMap = keyMap

	return m
}

// KeyMap returns a copy of the current key map in use.
func (m Model) KeyMap() KeyMap {
	return m.keyMap
}

// SelectableRows sets whether or not rows are selectable.  If set, adds a column
// in the front that acts as a checkbox and responds to controls if Focused.
func (m Model) SelectableRows(selectable bool) Model {
	const selectHeader = "[x]"
	m.selectableRows = selectable

	hasSelectColumn := m.columns[0].Key == columnKeySelect

	if hasSelectColumn != selectable {
		if selectable {
			m.columns = append([]Column{
				NewColumn(columnKeySelect, selectHeader, len(selectHeader)),
			}, m.columns...)
		} else {
			m.columns = m.columns[1:]
		}
	}

	m.recalculateWidth()

	return m
}

// HighlightedRow returns the full Row that's currently highlighted by the user.
func (m Model) HighlightedRow() Row {
	if len(m.rows) > 0 {
		return m.rows[m.rowCursorIndex]
	}

	// TODO: Better way to do this without pointers/nil?  Or should it be nil?
	return Row{}
}

// SelectedRows returns all rows that have been set as selected by the user.
func (m Model) SelectedRows() []Row {
	return m.selectedRows
}

// HighlightStyle sets a custom style to use when the row is being highlighted
// by the cursor.
func (m Model) HighlightStyle(style lipgloss.Style) Model {
	m.highlightStyle = style

	return m
}

// Focused allows the table to show highlighted rows and take in controls of
// up/down/space/etc to let the user navigate the table and interact with it.
func (m Model) Focused(focused bool) Model {
	m.focused = focused

	return m
}

// WithStaticFooter adds a footer that only displays the given text.
func (m Model) WithStaticFooter(footer string) Model {
	m.staticFooter = footer

	return m
}

// WithPageSize enables pagination using the given page size.
func (m Model) WithPageSize(pageSize int) Model {
	m.pageSize = pageSize

	return m
}

// WithNoPagination disable pagination in the table.
func (m Model) WithNoPagination() Model {
	m.pageSize = 0

	return m
}