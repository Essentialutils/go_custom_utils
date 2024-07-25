package go_custom_utils

import (
	"github.com/Delta456/box-cli-maker/v2"
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintInBox(title string, body string) {
	Box := box.New(box.Config{Px: 4,Py: 1, Type: "Round", Color: "Cyan", TitlePos: "Top"})
	Box.Print(title, body)
}

func PrintTable(data [][]string) {
	if len(data) == 0 {
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(data[0])

	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	table.SetAutoWrapText(true)

	for _, row := range data[1:] {
		table.Append(row)
	}

	table.Render()
}
