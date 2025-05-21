# Character Frequency Counter in Go

This project reads a text file and counts the occurrence of each alphabetic character (A-Z, a-z), ignoring digits and special characters.

The code follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) for organizing Go projects.

## 📁 Project Structure


- `cmd/main.go`: The application entry point, reads a filename from command line args.
- `internal/charcount/counter.go`: Implements character frequency counting logic.
- `internal/charcount/counter_test.go`: Unit tests for the counting logic.

## 🧠 Functionality

- Reads the contents of a specified text file.
- Counts occurrences of each alphabetic character, case-insensitive.
- Ignores digits, spaces, punctuation, and special characters.
- Prints frequency counts to the terminal.

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later installed

### Running the Application

Build and run with the filename as argument:

```bash
go run ./cmd sample.txt


🧪 Running Tests

go test ./internal/...

✅ Run Tests with Coverage

go test -cover ./internal/...

