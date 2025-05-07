package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/omar0ali/todos/core"
)

func LoadingData(filename string, rows any) {
	_, err := os.Stat(filename) //check the file
	if err != nil {
		f, err := os.Create(filename) // if it does not exit
		if err != nil {
			log.Panic(err)
		}
		defer f.Close() //closing file when we done
		f.Write([]byte("[]"))
	}
	var data []byte
	data, err = os.ReadFile(filename)
	if err != nil {
		log.Panic(err)
	}

	if len(data) == 0 { //additional check
		data = []byte("[]")
	}

	err = json.Unmarshal(data, rows)
	if err != nil {
		log.Panic(err)
	}

	if castedRows, ok := rows.(*[]core.TableCl); ok {
		if len(*castedRows) > 0 {
			core.LastId = (*castedRows)[len(*castedRows)-1].Id
		} else {
			core.LastId = 0 // or whatever default you prefer
		}
	} else {
		log.Println("Failed to cast rows to *[]core.TableCl")
	}
}

func SaveData(filename string, rows []core.TableCl) {
	data, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Panic(err)
	}
}
