package main

import (
	"fmt"
	"os"
)

func getFileName() string {
	return os.ExpandEnv("$HOME/.zsh_history")
}

func getNumberOfLines() int {
	var numLines int
	fmt.Print("Enter the number of last lines to display: ")
	fmt.Scanf("%d", &numLines)
	return numLines
}

func getOutputFileName() string {
	var outputFileName string
	fmt.Print("Enter the output file name (without extension): ")
	fmt.Scanf("%s", &outputFileName)
	return outputFileName + ".sh"
}
