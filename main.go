package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type ZshHistory struct {
	lines []string
}

type FileReader struct {
	fileName string
}

type FileWriter struct {
	fileName string
}

type Selection struct {
	numLines int
	indices  []int
}

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

func getFileName() string {
	return os.ExpandEnv("$HOME/.zsh_history")
}

func getNumberOfLines() int {
	var numLines int
	fmt.Print("Enter the number of last lines to display: ")
	fmt.Scanf("%d", &numLines)
	return numLines
}

func (fr *FileReader) readLines() (ZshHistory, error) {
	file, err := os.Open(fr.fileName)
	if err != nil {
		return ZshHistory{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isValidLine(line) {
			lines = append(lines, extractCommand(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return ZshHistory{}, err
	}

	return ZshHistory{lines: lines}, nil
}

func isValidLine(line string) bool {
	return strings.TrimSpace(line) != ""
}

func extractCommand(line string) string {
	if strings.HasPrefix(line, ":") {
		parts := strings.SplitN(line, ";", 2)
		if len(parts) == 2 {
			return parts[1]
		}
	}
	return line
}

func (zh *ZshHistory) displayLastLines(numLines int) {
	if numLines > len(zh.lines) {
		numLines = len(zh.lines)
	}

	start := len(zh.lines) - numLines
	for i, line := range zh.lines[start:] {
		fmt.Printf("%d: %s\n", i+start, line)
	}
}

func (zh *ZshHistory) selectLines(numLines int) []int {
	if numLines > len(zh.lines) {
		numLines = len(zh.lines)
	}

	start := len(zh.lines) - numLines
	items := make([]string, numLines)
	for i := 0; i < numLines; i++ {
		items[i] = zh.lines[start+i]
	}

	selectedIndices := promptUserSelection(items)

	// Adjust indices to match original line numbers
	for i := range selectedIndices {
		selectedIndices[i] += start
	}

	return selectedIndices
}

func promptUserSelection(items []string) []int {
	selected := []int{}
	prompt := &survey.MultiSelect{
		Message: "Select lines (use arrow keys and space to select, enter to confirm):",
		Options: items,
	}

	survey.AskOne(prompt, &selected)

	return selected
}

func getOutputFileName() string {
	var outputFileName string
	fmt.Print("Enter the output file name (without extension): ")
	fmt.Scanf("%s", &outputFileName)
	return outputFileName + ".sh"
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
