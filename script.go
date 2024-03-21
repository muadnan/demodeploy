package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	// Define the file name
	fileName := "CHANGELOG.md"
	msg := flag.String("msg", "", "Enter msg to write")
	flag.Parse()
	// Check if the file exists
	if _, err := os.Stat(fileName); err == nil {
		// Append the demo line to the end of the file
		err := appendDemoLine(fileName, *msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Demo changes documented in changelog.md")
	} else {
		// If the file doesn't exist, display an error message
		fmt.Printf("Error: File %s not found.\n", fileName)
	}
}

func appendDemoLine(fileName string, msg string) error {
	// Open the file in append mode
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the demo line to the file
	demoLine := "## Demo\n- Added demo line\nDemo complete\n" + msg
	if _, err := file.WriteString(demoLine); err != nil {
		return err
	}

	return nil
}
