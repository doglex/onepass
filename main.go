package main

import "flag"

func main() {
	// todo: list all history
	app := flag.Args()[0]
	Work(app)
}
