package main

import (
	"pack/solution"
)

func main() {
	credPath := "solution/cred.json"
	csvPath := "solution/data.csv"

	solution.CsvToCalendar(credPath, csvPath)
}
