package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a CSV file as an argument.")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			os.Exit(1)
		}
	}(file)

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header row
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var outputFile *os.File
	writer := csv.NewWriter(outputFile)
	counter := 0
	fileIndex := 0

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if counter%150 == 0 {
			if outputFile != nil {
				writer.Flush()
				err := outputFile.Close()
				if err != nil {
					fmt.Println("Error closing file:", err)
					os.Exit(1)
				}
			}

			outputFile, err = os.Create("import" + strconv.Itoa(fileIndex) + ".csv")
			if err != nil {
				fmt.Println("Error creating file:", err)
				os.Exit(1)
			}

			writer = csv.NewWriter(outputFile)
			fileIndex++
		}

		err = writer.Write(record)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		counter++
	}

	writer.Flush()
	if outputFile != nil {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			os.Exit(1)
		}
	}
}
