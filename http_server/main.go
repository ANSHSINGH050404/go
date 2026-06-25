package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("messages.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// Read the contents of the file
	buffer := make([]byte, 8)

	currentLine := ""

	for {
		n, err := file.Read(buffer)

		if n > 0 {
			data := string(buffer[:n])
			// println(data)

			parts := strings.Split(data, "\n")
			// println(parts)

			// Process all except last
			for i := 0; i < len(parts)-1; i++ {
				line := currentLine + parts[i]
				fmt.Println("read:", line)
				currentLine = ""
			}

			// Last part (possibly incomplete)
			currentLine += parts[len(parts)-1]

		}

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	// Print leftover
	if currentLine != "" {
		fmt.Println("read:", currentLine)
	}

}
