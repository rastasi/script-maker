package main

import (
	"bufio"
	"os"
	"strings"
)

type FileReader struct {
	fileName string
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
