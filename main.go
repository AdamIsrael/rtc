package main

import "github.com/adamisrael/rtc/cmd"

/*
+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

| status | min | max | avg |
| 0	  | 0.01 | 0.02 | 0.015 |
| 1	  | 0.02 | 0.03 | 0.025 |
*/

// var (
// 	colTitleStatus  = "Status"
// 	colTitleMin     = "Min"
// 	colTitleMax     = "Max"
// 	colTitleAverage = "Avg"
// 	colTitleCount   = "Count"
// 	rowHeader       = table.Row{colTitleStatus, colTitleCount, colTitleMin, colTitleMax, colTitleAverage}
// 	row1            = table.Row{0, 80, "0.01", "0.02", "0.015"}
// 	row2            = table.Row{1, 20, "0.02", "0.03", "0.025"}
// )

func main() {
	cmd.Execute()

	// t := table.NewWriter()

	// t.AppendHeader(rowHeader)

	// t.AppendRow(row1)
	// t.AppendRow(row2)

	// fmt.Println(t.Render())
}
