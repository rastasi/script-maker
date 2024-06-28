package main

import (
	"fmt"
	"os"
)

type FileWriter struct {
	fileName string
}

func (fw *FileWriter) saveSelectedLines(history ZshHistory, indices []int) {
	outputFile, err := os.Create(fw.fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	for _, index := range indices {
		if isIndexInRange(index, history.lines) {
			writeLineToFile(outputFile, history.lines[index])
		}
	}

	makeFileExecutable(fw.fileName)
}

func isIndexInRange(index int, lines []string) bool {
	return index >= 0 && index < len(lines)
}

func writeLineToFile(outputFile *os.File, line string) {
	_, err := outputFile.WriteString(line + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func makeFileExecutable(fileName string) {
	err := os.Chmod(fileName, 0755)
	if err != nil {
		fmt.Println("Error making file executable:", err)
		return
	}

	fmt.Println("Script saved and made executable:", fileName)
}
