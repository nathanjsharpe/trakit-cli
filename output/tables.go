package output

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintSimpleTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()
}

func PrintResourceTable(data [][]string, headers []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(70)
	table.AppendBulk(data)
	table.SetHeader(headers)
	table.Render()
}

func PrintBorderlessTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetColWidth(75)
	table.AppendBulk(data)
	table.Render()
}
