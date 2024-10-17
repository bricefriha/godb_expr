package services

import (
	"fmt"
	"regexp"
	"strings"
)

func Execute(query string) string {
	var fields = strings.Fields(query)

	switch fields[0] {
	case "INSERT":
		if fields[1] != "INTO" {
			throwSyntaxError()
			return "Wrong Syntax"
		}
		// get address
		var addr []string = strings.Split(fields[2], ".")

		println(addr[0])
		println(addr[1])

		re := regexp.MustCompile(`\(([^)]+)\)`)
		columnsRaw := re.FindStringSubmatch(query)[1]

		if len(columnsRaw) < 3 {
			throwSyntaxError()
			return "Wrong Syntax"
		}

		var columns []string = strings.Split(columnsRaw, ",")

		println("columns:")
		for i := 0; i < len(columns); i++ {
			println(columns[i])
		}

	}

	return "fail"
}
func throwSyntaxError() {
	fmt.Println("Wrong Syntax")
}
