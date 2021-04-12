package display

import (
	"fmt"
	"log"
	"time"

	"github.com/Eldius/display-metrics-go/metrics"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	dateFormatPattern = "2006-01-02 15:04:05-0700"
	fetchTimeInterval = "15s"
)

func Display() {
	for {
		_, err := metrics.GetSummary()
		if err == nil {
			break
		}
		log.Println("Waiting metrics become available...")
		time.Sleep(10 * time.Second)
	}
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	go fetchMetrics()
	metrics.GetSummary()
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}

}

func fetchMetrics() {
	d, _ := time.ParseDuration(fetchTimeInterval)
	for {
		m, _ := metrics.GetSummary()
		printSummary(m)
		time.Sleep(d)
	}
}

func printSummary(m *metrics.SummaryResponse) {
	np := widgets.NewParagraph()
	np.Text = fmt.Sprintf("nodes\n%d", m.Data.Nodes)
	np.SetRect(0, 0, 7, 4)
	np.BorderStyle.Fg = ui.ColorGreen
	np.TextStyle.Fg = ui.ColorYellow

	pc := widgets.NewParagraph()
	pc.Text = fmt.Sprintf("pods\n%d", m.Data.Pods)
	pc.SetRect(8, 0, 15, 4)
	pc.BorderStyle.Fg = ui.ColorGreen
	pc.TextStyle.Fg = ui.ColorYellow

	cc := widgets.NewParagraph()
	cc.Text = fmt.Sprintf("containers\n%d", m.Data.Containers)
	cc.SetRect(16, 0, 28, 4)
	cc.BorderStyle.Fg = ui.ColorGreen
	cc.TextStyle.Fg = ui.ColorYellow

	mu := widgets.NewParagraph()
	mu.Text = fmt.Sprintf("memory (%%)\n%3.4f", m.Data.Memory)
	mu.SetRect(0, 4, 13, 8)
	mu.BorderStyle.Fg = ui.ColorGreen
	mu.TextStyle.Fg = ui.ColorYellow

	cu := widgets.NewParagraph()
	cu.Text = fmt.Sprintf("cpu (%%)\n%3.4f", m.Data.CPU)
	cu.SetRect(14, 4, 28, 8)
	cu.BorderStyle.Fg = ui.ColorGreen
	cu.TextStyle.Fg = ui.ColorYellow

	updated := widgets.NewParagraph()
	updated.Text = time.Now().Format(dateFormatPattern)
	updated.SetRect(0, 8, 28, 11)
	updated.BorderStyle.Fg = ui.ColorGreen
	updated.TextStyle.Fg = ui.ColorYellow

	ui.Render(np)
	ui.Render(pc)
	ui.Render(cc)
	ui.Render(mu)
	ui.Render(cu)
	ui.Render(updated)
}
