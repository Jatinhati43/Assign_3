package main

import (
	"fmt"
	"log"
	"path/filepath"

	"ASSIGN_3/internal/charcount"
)

func main() {
	filePath := filepath.Join(".", "sample.txt")

	counts, err := charcount.CountCharactersFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Character Frequencies:")
	for ch, count := range counts {
		fmt.Printf("%c: %d\n", ch, count)
	}
}
