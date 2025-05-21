package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Jatinhati43/Assign_3/internal/charcount"
)

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
