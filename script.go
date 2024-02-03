package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	// Define the file name
	fileName := "CHANGELOG.md"

	// Check if the file exists
	if _, err := os.Stat(fileName); err == nil {
		// Append the demo line to the end of the file
		err := appendDemoLine(fileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Demo changes documented for golang in changelog.md")
	} else {
		// If the file doesn't exist, display an error message
		fmt.Printf("Error: File %s not found.\n", fileName)
	}
}

func appendDemoLine(fileName string) error {
	// Open the file in append mode
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the demo line to the file
	demoLine := "## Demo\n- Added demo line for golang\nDemo complete\n"
	if _, err := file.WriteString(demoLine); err != nil {
		return err
	}

	return nil
}
