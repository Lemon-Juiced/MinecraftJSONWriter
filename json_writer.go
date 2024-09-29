package main

import (
	"fmt"
	"os"
)

/**
 * Writes the item model JSON files for the given resource name and namespace.
 *
 * @param filename The name of the file.
 * @param content The content to write to the file.
 */
func writeJsonFile(filename string, content string) {
	// Write the JSON file.
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
