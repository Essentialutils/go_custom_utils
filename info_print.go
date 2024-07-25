package go_custom_utils

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"slices"
)

type InfoPrintData struct {
	ServerRunningPort string
	ServerRunningMode string
	OpenApiPath       string
	Developer         string
	Website           string
	Version           string
}

func InfoPrint(v InfoPrintData) {
	ips := GetWifiIP()

	data := [][]string{
		{"CATEGORY", "SUB CATEGORY", "VALUES"},
		{"Server", "mode", v.ServerRunningMode},
		{"Server", "host", "http://localhost" + v.ServerRunningPort},
		{"Server", "host", "http://[::1]" + v.ServerRunningPort},
	}

	index := slices.IndexFunc(ips, func(ipMap map[string]string) bool {
		return ipMap["name"] == "Wi-Fi"
	})

	if index != -1 {
		ipMap := ips[index]
		data = append(data, []string{"Server", "host", "http://" + ipMap["ip"] + v.ServerRunningPort})
	}

	data = append(data, []string{"Server", "Api Collection Path", v.OpenApiPath})

	for _, ipMap := range ips {
		data = append(data, []string{"Connection", ipMap["name"], ipMap["ip"]})
	}

	var buffer bytes.Buffer

	table := tablewriter.NewWriter(&buffer)
	table.SetHeader(data[0])

	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	table.SetAutoWrapText(true)

	for _, row := range data[1:] {
		table.Append(row)
	}

	table.Render()

	tableText := buffer.String()

	PrintInBox(
		"Info",
		fmt.Sprintf(`Developer : %s
Website   : %s
Version   : %s

%s`,
			v.Developer,
			v.Website,
			v.Version,
			tableText,
		),
	)
}
