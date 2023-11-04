package tui

// import (
// 	"fmt"
// 	"io"
// 	// "os"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// const listHeight = 14

// var (
// 	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
// 	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
// 	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
// 	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
// 	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
// 	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
// )

// type command string
// type commandDelegate struct{}

// type WizardModel struct {

// }

// func (d commandDelegate) Height() int  { return 1 }
// func (d commandDelegate) Spacing() int { return 0 }
// func (d commandDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
// 	i, ok := listItem.(item)
// 	if !ok {
// 		return
// 	}

// 	str := fmt.Sprintf("%d. %s", index+1, i)

// 	fn := itemStyle.Render
// 	if index == m.Index() {
// 		fn = func(s ...string) string {
// 			return selectedItemStyle.Render("> " + strings.Join(s, " "))
// 		}
// 	}

// 	fmt.Fprint(w, fn(str))
// }

// func (w WizardModel) Init() tea.Cmd {

// }

// func (w WizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

// }

// func (w WizardModel) View() string {

// }
