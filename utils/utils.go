package utils

import (
	"fmt"
	"log"
	"os"

	"encoding/json"
)

func LoadingData(filename string, rows any) {
	_, err := os.Stat(filename) //check the file
	if err != nil {
		f, err := os.Create(filename) // if it does not exit
		if err != nil {
			log.Panic(err)
		}
		defer f.Close() //closing file when we done
	}
	var data []byte
	data, err = os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &rows)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(rows)
}

func SaveData(filename string, rows any) {
	data, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Panic(err)
	}
}
