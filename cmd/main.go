package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Jatinhati43/Assign_3/internal/charcount"
)

// main reads a text file path from the command line argument,
// reads the file line by line, counts the occurrence of each
// alphabetic character (ignoring digits and special characters)
// using the CountChars function from the internal charcount package,
// and then prints the aggregated counts to the console.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file-path>")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	charCounts := make(map[rune]int)

	for scanner.Scan() {
		line := scanner.Text()
		lineCounts := charcount.CountChars(line)
		for k, v := range lineCounts {
			charCounts[k] += v
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println("Character counts (letters only):")
	for ch, count := range charCounts {
		fmt.Printf("%c : %d\n", ch, count)
	}
}
