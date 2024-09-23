package output

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

type CommandExecution struct {
	ExitCode int
	Runtime  time.Duration
}

var (
	colTitleStatus  = "Status"
	colTitleMin     = "Min"
	colTitleMax     = "Max"
	colTitleAverage = "Avg"
	colTitleCount   = "Count"
	rowHeader       = table.Row{colTitleStatus, colTitleCount, colTitleMin, colTitleMax, colTitleAverage}
	row1            = table.Row{0, 80, "0.01", "0.02", "0.015"}
	row2            = table.Row{1, 20, "0.02", "0.03", "0.025"}
)

func DisplayTable(commandExecutions []CommandExecution) {
	t := table.NewWriter()

	t.AppendHeader(rowHeader)

	//
	for _, commandExecution := range commandExecutions {

		t.AppendRow(table.Row{commandExecution.ExitCode, commandExecution.Runtime})
	}
	// t.AppendRow(row1)
	// t.AppendRow(row2)

	fmt.Println(t.Render())
}
