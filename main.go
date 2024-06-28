package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fileName := getFileName()
	fmt.Println("Reading from file:", fileName)

	n := getNumberOfLines()

	lines, err := readLinesFromFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	displayLastLines(lines, n)

	selectedIndices := selectLines(lines, n)
	outputFileName := getOutputFileName()
	saveSelectedLines(lines, selectedIndices, outputFileName)
}

func getFileName() string {
	return os.ExpandEnv("$HOME/.zsh_history")
}

func getNumberOfLines() int {
	var n int
	fmt.Print("Enter the number of last lines to display: ")
	fmt.Scanf("%d", &n)
	return n
}

func readLinesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if validLine(line) {
			lines = append(lines, extractCommand(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func validLine(line string) bool {
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

func displayLastLines(lines []string, n int) {
	if n > len(lines) {
		n = len(lines)
	}

	start := len(lines) - n
	for i, line := range lines[start:] {
		fmt.Printf("%d: %s\n", i+start, line)
	}
}

func selectLines(lines []string, n int) []int {
	if n > len(lines) {
		n = len(lines)
	}

	start := len(lines) - n
	items := make([]string, n)
	for i := 0; i < n; i++ {
		items[i] = lines[start+i]
	}

	selected := []int{}
	prompt := &survey.MultiSelect{
		Message: "Select lines (use arrow keys and space to select, enter to confirm):",
		Options: items,
	}

	survey.AskOne(prompt, &selected)

	// Adjust indices to match original line numbers
	for i := range selected {
		selected[i] += start
	}

	return selected
}

func getOutputFileName() string {
	var outputFileName string
	fmt.Print("Enter the output file name (without extension): ")
	fmt.Scanf("%s", &outputFileName)
	return outputFileName + ".sh"
}

func saveSelectedLines(lines []string, indices []int, outputFileName string) {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	for _, index := range indices {
		if indexInRange(index, lines) {
			writeLineToFile(outputFile, lines[index])
		}
	}

	makeFileExecutable(outputFileName)
}

func indexInRange(index int, lines []string) bool {
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
