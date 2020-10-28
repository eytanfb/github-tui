package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shurcooL/githubv4"
	"github.com/skanehira/ght/config"
	"github.com/skanehira/ght/github"
)

type Label struct {
	Name string
}

type LabelsUI struct {
	updater func(f func())
	*tview.Table
}

func NewLabelsUI(updater func(f func())) *LabelsUI {
	ui := &LabelsUI{
		Table:   tview.NewTable().SetSelectable(true, false).Select(0, 0).SetFixed(0, 0),
		updater: updater,
	}
	ui.SetBorder(true).SetTitle("label list").SetTitleAlign(tview.AlignLeft)
	ui.updateLabelList()
	return ui
}

func (ui *LabelsUI) updateLabelList() {
	table := ui.Clear()

	ui.updater(func() {
		v := map[string]interface{}{
			"owner":  githubv4.String(config.GitHub.Owner),
			"name":   githubv4.String(config.GitHub.Repo),
			"first":  githubv4.Int(100),
			"cursor": (*githubv4.String)(nil),
		}
		resp, err := github.GetRepoLabels(v)
		if err != nil {
			log.Println(err)
			return
		}

		labels := make([]Label, len(resp.Nodes))

		for i, l := range resp.Nodes {
			name := string(l.Name)
			labels[i] = Label{
				Name: name,
			}
			table.SetCell(i, 1, tview.NewTableCell(name).
				SetTextColor(tcell.ColorWhite).SetExpansion(1))
		}

		ui.ScrollToBeginning()
	})
}