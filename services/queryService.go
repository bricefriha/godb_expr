package services

import (
	"fmt"
	"strings"
)

func Execute(query string) string {
	var fields = strings.Fields(query)

	switch fields[0] {
	case "INSERT":
		if fields[1] != "INTO" {
			throwSyntaxError()
		}
		var addr []string = strings.Split(fields[2], ".")

		println(addr[0])
		println(addr[1])

	}

	return "fail"
}
func throwSyntaxError() {
	fmt.Println("Wrong Syntax")
}
