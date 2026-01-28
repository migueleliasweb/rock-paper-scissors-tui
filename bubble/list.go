package bubble

import "github.com/charmbracelet/bubbles/list"

// ItemWithDeactivation Auguments the functionality from the upstream list bubble
// by adding the ability to handle deactivated items.
type ItemWithDeactivation struct {
	// TitleItem Item's tittle
	TitleItem string

	// DescItem Item's description
	DescItem string

	// Deactivated Whether the item should be renderer as `deactivated`
	Deactivated bool
}

func (i ItemWithDeactivation) Title() string {
	return i.TitleItem
}

func (i ItemWithDeactivation) Description() string {
	return i.DescItem
}

func (i ItemWithDeactivation) FilterValue() string {
	return i.TitleItem
}

var _ list.Item = &ItemWithDeactivation{}

// SimpleItem implements the list.Item interface
type SimpleItem struct {
	// Title Item Item's tittle
	TitleItem string

	// Desc Item Item's description
	DescItem string
}

func (i SimpleItem) Title() string       { return i.TitleItem }
func (i SimpleItem) Description() string { return i.DescItem }
func (i SimpleItem) FilterValue() string { return i.TitleItem }

var _ list.Item = &SimpleItem{}

func delegateItemWithDeactivation() {

}

// func newItemDelegate() list.DefaultDelegate {
// 	d := list.NewDefaultDelegate()

// 	// Define a specific style for disabled text (dim/grey)
// 	disabledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Grey

// 	// Override the Render function
// 	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
// 		return nil
// 	}

// 	// This is where the visual magic happens
// 	d.Render = func(w io.Writer, m list.Model, index int, listItem list.Item) {
// 		i, ok := listItem.(ItemWithDeactivation)
// 		if !ok {
// 			return
// 		}

// 		// 1. If it's the currently selected item...
// 		if index == m.Index() {
// 			// If it's disabled, we render it differently even if selected
// 			if i.disabled {
// 				// Render a "cannot select" indicator (like a crossed circle or just dim text)
// 				fmt.Fprint(w, disabledStyle.Render("> "+i.Title()+" (Unavailable)"))
// 				return
// 			}
// 			// Otherwise render the standard selected style
// 			// (We invoke the default logic for ease, or write custom logic here)
// 			d.Styles.SelectedTitle.Render(w) // Helper to access default styles
// 		}

// 		// 2. Standard rendering for unselected items
// 		if i.disabled {
// 			fmt.Fprint(w, disabledStyle.Render("  "+i.Title()))
// 			fmt.Fprint(w, disabledStyle.Render("\n    "+i.Description()))
// 		} else {
// 			// Render normal text
// 			fmt.Fprint(w, "  "+i.Title())
// 			fmt.Fprint(w, "\n    "+lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Render(i.Description()))
// 		}
// 	}

// 	// Set height of each row
// 	d.SetHeight(2)
// 	return d
// }

// // --- 3. Main Model Logic ---

// type model struct {
// 	list list.Model
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		if msg.String() == "enter" {
// 			// LOGIC GUARD: Prevent action if item is disabled
// 			selectedItem := m.list.SelectedItem().(ItemWithDeactivation)
// 			if selectedItem.disabled {
// 				// Do nothing, or show an error message
// 				return m, nil
// 			}
// 		}
// 		if msg.String() == "ctrl+c" {
// 			return m, tea.Quit
// 		}
// 	}

// 	var cmd tea.Cmd
// 	m.list, cmd = m.list.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	return "\n" + m.list.View()
// }

// func main() {
// 	items := []list.Item{
// 		ItemWithDeactivation{title: "Production DB", desc: "US-East-1", disabled: false},
// 		ItemWithDeactivation{title: "Staging DB", desc: "US-West-2", disabled: false},
// 		ItemWithDeactivation{title: "Legacy DB", desc: "Offline for maintenance", disabled: true}, // Disabled!
// 		ItemWithDeactivation{title: "Analytics", desc: "EU-Central-1", disabled: false},
// 	}

// 	// Initialize list with our custom delegate
// 	l := list.New(items, newItemDelegate(), 20, 14)
// 	l.Title = "Server Status"
// 	l.SetShowStatusBar(false)

// 	p := tea.NewProgram(model{list: l})
// 	if _, err := p.Run(); err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// }
