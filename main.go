package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.ExpandEnv("$HOME/.zsh_history")
	fmt.Println("Reading from file:", fileName)

	var n int
	fmt.Print("Enter the number of last lines to display: ")
	fmt.Scanf("%d", &n)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			if strings.HasPrefix(line, ":") {
				parts := strings.SplitN(line, ";", 2)
				if len(parts) == 2 {
					lines = append(lines, parts[1])
				}
			} else {
				lines = append(lines, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if n > len(lines) {
		n = len(lines)
	}

	start := len(lines) - n
	for i, line := range lines[start:] {
		fmt.Printf("%d: %s\n", i+start, line)
	}

	var indices string
	fmt.Print("Enter the line numbers to save (comma-separated): ")
	fmt.Scanf("%s", &indices)

	fmt.Print("Enter the output file name (without extension): ")
	var outputFileName string
	fmt.Scanf("%s", &outputFileName)
	outputFileName += ".sh"

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	indicesList := strings.Split(indices, ",")
	for _, indexStr := range indicesList {
		index, err := strconv.Atoi(strings.TrimSpace(indexStr))
		if err != nil {
			fmt.Println("Invalid index:", indexStr)
			continue
		}
		if index >= start && index < len(lines) {
			_, err := outputFile.WriteString(lines[index] + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		} else {
			fmt.Println("Index out of range:", index)
		}
	}

	err = outputFile.Chmod(0755)
	if err != nil {
		fmt.Println("Error making file executable:", err)
		return
	}

	fmt.Println("Script saved and made executable:", outputFileName)
}
