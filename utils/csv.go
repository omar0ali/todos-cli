package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

func ReadCSVFile(filename string) (*csv.Reader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bytes.NewReader(data)) // parsing
	return reader, nil
}

func DisplayRecords(reader *csv.Reader) {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	tabWriter.Write([]byte("ID\tTitle\tDscription\tCreated\tStatus"))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading CSV data:", err)
			break
		}
		fmt.Println("[DEBUG] [OUTPUT] [CSV] ", record)
		if len(record) < 4 {
			fmt.Println("Skipping malformed record:", record)
			continue
		}

		id := record[0]
		title := record[1]
		description := record[2]
		status := record[3]
		fmt.Fprintf(tabWriter, "%s\t%s\t%s\t%s\n", id, title, description, status)
	}
	tabWriter.Flush()
}
