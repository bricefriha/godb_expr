package main

import (
	"godb/services"
	"os"
)

func main() {
	if len(os.Args) > 0 && os.Args[1][0] == '-' {
		action := os.Args[1]

		switch action {
		case "-exc":
			println(services.Execute(os.Args[2]))
		default:
			return

		}

	}
}
