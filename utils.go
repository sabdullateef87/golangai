package main

import (
	"encoding/json"
	"fmt"
)

// PrintStruct prints any struct in a readable format.
// Use format "json" for pretty JSON, or anything else for Go's default %+v output.
func PrintStruct[T any](s T, format string) {
	if format == "json" {
		data, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println(string(data))
	} else {
		fmt.Printf("%+v\n", s)
	}
}

