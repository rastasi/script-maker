package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type ZshHistory struct {
	lines []string
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
