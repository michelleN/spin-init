package main

import (
	"fmt"
	"os"
)

func main() {

	if fileExists("spin.toml") || fileExists("Spin.toml") {
		fmt.Println("Error: spin.toml or Spin.toml already exists.")
		return
	}

	tomlContent := `
spin_manifest_version = 2

[application]
name = ""
version = "0.1.0"
authors = ["YOUR NAME <youremail@domain>"]
description = ""

`

	file, err := os.Create("spin.toml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(tomlContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("spin.toml file created successfully.")

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
