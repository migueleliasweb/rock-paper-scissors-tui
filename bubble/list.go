package bubble

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

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

type delegate struct {
	list.DefaultDelegate
}

func (d delegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(ItemWithDeactivation)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, listItem)
		return
	}

	if item.Deactivated {
		dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		title := item.Title()
		desc := item.Description()

		if index == m.Index() {
			fmt.Fprint(w, dimStyle.Render("> "+title+" (Unavailable)"))
		} else {
			fmt.Fprint(w, dimStyle.Render("  "+title))
		}
		fmt.Fprint(w, "\n"+dimStyle.Render("    "+desc))
		return
	}

	d.DefaultDelegate.Render(w, m, index, listItem)
}

func DelegateItemWithDeactivation() list.ItemDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(
		lipgloss.Color("#7D56F4"),
	).BorderForeground(lipgloss.Color("#7D56F4"))

	d.Styles.SelectedDesc = d.Styles.SelectedDesc.Foreground(
		lipgloss.Color("#505164"),
	).BorderForeground(lipgloss.Color("#7D56F4"))

	return delegate{DefaultDelegate: d}
}
