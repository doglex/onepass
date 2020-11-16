package main

import "flag"

func main() {
	if ListApp || len(flag.Args()) == 0 {
		HistoryList()
	} else {
		app := flag.Args()[0]
		Work(app)
	}
}
