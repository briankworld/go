package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file", err)
		return
	}
	defer inputFile.Close()

	outFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file", err)
		return
	}
	defer outFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outFile)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error reading file", err)
			return
		}

		if _, err := writer.WriteString(strings.ToUpper(line)); err != nil {
			fmt.Println("error writing to file", err)
			return
		}
	}
	writer.Flush()
}
