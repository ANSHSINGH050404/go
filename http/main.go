package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("file opened successfully %v\n", file)

	// This variable persists between loop iterations
	var currentLine string
	data := make([]byte, 8)

	fmt.Println("data:", data)

	for {
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", n, data[:n])

		// 1. Convert the current chunk to a string
		chunk := string(data[:n])

		fmt.Printf("chunk read: %q\n", chunk)

		// 2. Split the chunk on newlines
		parts := strings.Split(chunk, "\n")

		fmt.Printf("parts: %q\n", parts)

		// 3. Process all parts except the last one
		// These represent completed lines (because a newline followed them)
		for i := 0; i < len(parts)-1; i++ {
			fullLine := currentLine + parts[i]
			fmt.Printf("read: %s\n", fullLine)
			currentLine = "" // Reset for the next line
		}

		

		// 4. Add the last part to currentLine
		// It's either the start of a new line or a continuation of one
		currentLine += parts[len(parts)-1]

		fmt.Printf("currentLine after processing: %q\n", currentLine)
	}

	// 5. After the loop, print any remaining text (the final line)
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}