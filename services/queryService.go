package services

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Execute(query string) string {
	var fields = strings.Fields(query)

	switch fields[0] {
	case "CREATE":
		if fields[1] != "TABLE" {
			throwSyntaxError()
			return "Wrong Syntax"
		}
		addr := strings.Split(fields[2], ".")
		if len(addr) < 2 {
			return "database unspecified"
		}
		CreateTable(addr[1], fmt.Sprintf("exampleDocs/%s.json", addr[0]))
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

		// Extract columns
		var columns []string = strings.Split(columnsRaw, ",")

		// Extract values
		var valre = regexp.MustCompile(`(?i)VALUES\s*\((.*?)\)`)
		var valRaw = valre.FindStringSubmatch(query)[1]

		if len(valRaw) < 3 {
			throwSyntaxError()
			return "Wrong Syntax"
		}

		var vals []string = strings.Split(valRaw, ",")
		var jsonData string = ``

		for i := 0; i <= len(columns); i++ {
			if i == 0 {
				jsonData += "{"
			}
			if i == len(columns) {
				jsonData += "}"
			} else {
				jsonData += fmt.Sprintf(`"%s": %s`, columns[i], legilimiseValue(vals[i]))
				if i < len(columns)-1 {
					jsonData += ","
				}
			}
		}
		// Execute insert
		Insert(jsonData, addr[0], addr[1])
		return jsonData

	}

	return "fail"
}
func throwSyntaxError() {
	fmt.Println("Wrong Syntax")
}
func legilimiseValue(input string) string {
	// Check if the input is a boolean
	if strings.ToLower(input) == "true" || strings.ToLower(input) == "false" {
		return input
	}

	// Check if the input is a number
	if _, err := strconv.ParseFloat(input, 64); err == nil {
		return input
	}

	// If it's neither a boolean nor a number, add quotes
	return fmt.Sprintf(`"%s"`, input)
}
