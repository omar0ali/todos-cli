// Package utils
package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/omar0ali/todos/core"
)

func GetCurrentDir(verbose bool) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		log.Printf("[DIRECTORY: %s]\n", cwd)
	}
	return cwd
}

func LoadingData(filename string, rows any) {
	path := GetCurrentDir(false)
	filePath := filepath.Join(path, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()
		_, err = f.Write([]byte("[]"))
		if err != nil {
			log.Panic(err)
		}
	}

	var data []byte
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	if len(data) == 0 { // additional check
		data = []byte("[]")
	}

	err = json.Unmarshal(data, rows)
	if err != nil {
		log.Panic(err)
	}

	if castedRows, ok := rows.(*[]core.TableCl); ok {
		if len(*castedRows) > 0 {
			core.LastID = (*castedRows)[len(*castedRows)-1].ID
		} else {
			core.LastID = 0 // or whatever default you prefer
		}
	} else {
		log.Println("Failed to cast rows to *[]core.TableCl")
	}
}

func SaveData(filename string, rows []core.TableCl) {
	path := GetCurrentDir(false)
	filePath := filepath.Join(path, filename)

	data, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Panic(err)
	}
}
