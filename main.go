package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Data struct {
    Column1 string
    Column2 string
    Column3 string
}

func main() {
    // Open the CSV file
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Read the CSV data
    reader := csv.NewReader(file)

    // Skip the header row
    _, err = reader.Read()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    lines, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Parse the CSV data into a slice of structs
    var data []Data
    for _, line := range lines {
        record := Data{
            Column1: line[0],
            Column2: line[1],
            Column3: line[2],
        }
        data = append(data, record)
    }

    // Output the data in a formatted way
    fmt.Printf("%-15s%-15s%-15s\n", "Column 1", "Column 2", "Column 3")
    for _, record := range data {
        fmt.Printf("%-15s%-15s%-15s\n", record.Column1, record.Column2, record.Column3)
    }
}
