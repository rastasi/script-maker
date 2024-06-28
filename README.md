# Script Maker
Create scripts from zsh_history (AI generated)

## Prompt

Please write a Go program according to the following specifications to comply with Clean Code principles, use separate structs for file reading and file writing, and organize the code into separate files:

Specification:
The program reads commands from the user's .zsh_history file. The program should:

Ask the user to specify a number that determines how many lines from the end of the file to display.
Display the specified number of lines from the end of the .zsh_history file.
Allow the user to select lines using the up and down arrow keys and the spacebar to mark lines they want to save to a new file.
Prompt the user to enter the name of the new file.
Save the selected lines to the new file.
Make the new file executable.
File Structure:
main.go: Contains the main program logic.
file_reader.go: Responsible for file reading, containing a FileReader struct and its methods.
file_writer.go: Responsible for file writing, containing a FileWriter struct and its methods.
zsh_history.go: Contains a ZshHistory struct and its methods for handling the history commands.
selection.go: Contains the logic for user selection, including prompting the user and handling their input.
utils.go: Contains utility functions, such as those for getting the file name and the number of lines to display.
Main Requirements:
The main.go file should contain the main logic that orchestrates the program's flow:

Reading the file.
Displaying lines.
Handling user selection.
Saving selected lines to a new file.
The file_reader.go file should:

Define a FileReader struct with a method to read lines from a file.
Implement logic to open the file, read its contents, and handle errors.
The file_writer.go file should:

Define a FileWriter struct with a method to write lines to a file.
Implement logic to create the file, write selected lines, make the file executable, and handle errors.
The zsh_history.go file should:

Define a ZshHistory struct with methods to manage and display history lines.
Implement logic to process and extract relevant commands from the history file.
The selection.go file should:

Implement logic for user selection of lines using the up and down arrow keys and the spacebar.
Use a suitable library (e.g., github.com/AlecAivazis/survey/v2) to handle interactive prompts.
The utils.go file should:

Contain helper functions for common tasks such as getting the file name, prompting for the number of lines, and getting the output file name.
