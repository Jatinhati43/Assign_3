// Description: This program reads a text file (sample.txt),
//              counts the frequency of each alphabet character (ignoring digits and special characters),


package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Jatinhati43/Assign_3/internal/charcount"
)

func main() {
	filePath := filepath.Join(".", "sample.txt")

	counts, err := charcount.CountCharacters(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Character Frequencies:")
	for ch, count := range counts {
		fmt.Printf("%c: %d\n", ch, count)
	}
}
