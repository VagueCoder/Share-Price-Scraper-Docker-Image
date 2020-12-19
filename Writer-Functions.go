package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

func writeJSON(collection *Collection) {
	file, err := json.MarshalIndent(collection, "", "\t")
	if err != nil {
		fmt.Printf("\nRecord Failed to Marshal: %+v\n", collection)
	}

	var filename string = collection.Data.Name
	unwanted := []string{"/", "\\", "<", ">", ":", "\"", "|", "?", "*"}
	for _, chr := range unwanted {
		filename = strings.ReplaceAll(filename, chr, "_")
	}

	err = ioutil.WriteFile("DataStore/" + filename + ".json", file, 0644)
	if err != nil {
		fmt.Printf("\nRecord Failed to Write to JSON File: %+v\n", collection)
	}
}

func writeStat(flag bool, line string) {
	if flag {
		err := ioutil.WriteFile("DataStore/stats.txt", []byte(""), 0644)
		if err != nil {
			fmt.Printf("\nWriter Error: Couldn't Erase File stats.txt")
		}
	}

	file, err := os.OpenFile("DataStore/stats.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("\nOS Error: File stats.txt Creation Failed\n")
	}
	defer file.Close()

	if _, err := file.WriteString(line); err != nil {
		fmt.Printf("\nWriter Error: Writing on File stats.txt Failed. Context: %d\n", line)
	}
}