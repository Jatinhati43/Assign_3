package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jatinhati43/Assign_3/internal/charcount"
	"github.com/Jatinhati43/Assign_3/internal/File_reader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file-path>")
		return
	}

	lines, err := reader.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	charCounts := make(map[rune]int)
	for _, line := range lines {
		lineCounts := charcount.CountChars(line)
		for k, v := range lineCounts {
			charCounts[k] += v
		}
	}

	fmt.Println("Character counts (letters only):")
	for ch, count := range charCounts {
		fmt.Printf("%c : %d\n", ch, count)
	}
}
