package main

import (
	"pack/solution"
)

func main() {
	// fmt.Println(solution.Towers())
	credPath := "solution/cred.json"
	csvPath := "solution/data.csv"

	solution.CsvToCalendar(credPath, csvPath)
}
