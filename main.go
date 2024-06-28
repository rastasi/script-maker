package main

import (
	"fmt"
)

func main() {
	fileReader := FileReader{fileName: getFileName()}
	fmt.Println("Reading from file:", fileReader.fileName)

	history, err := fileReader.readLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	selection := Selection{numLines: getNumberOfLines()}
	history.displayLastLines(selection.numLines)

	selection.indices = history.selectLines(selection.numLines)
	outputFileName := getOutputFileName()
	fileWriter := FileWriter{fileName: outputFileName}
	fileWriter.saveSelectedLines(history, selection.indices)
}
