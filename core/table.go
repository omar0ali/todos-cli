package core

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

var (
	LastId uint = 0
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type TableCl struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Status      Status    `json:"status"`
}

func NewRow(title string, desc string, status Status) *TableCl {
	LastId++
	return &TableCl{
		Id:          LastId,
		Title:       title,
		Description: desc,
		Created:     time.Now(),
		Status:      status,
	}
}

func DisplayRecords(rows []TableCl, verbose bool) {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	tabWriter.Write([]byte("ID\tTitle\tDscription\tCreated\tStatus\n"))
	for i := 0; i < len(rows); i++ {
		fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%s\t%s\n", rows[i].Id, rows[i].Title, rows[i].Description, rows[i].Created, rows[i].Status)
	}
	tabWriter.Flush()
}

func DisplayRecordsStatus(rows []TableCl, verbose bool, status Status) {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	tabWriter.Write([]byte("ID\tTitle\tDscription\tCreated\tStatus\n"))
	for i := 0; i < len(rows); i++ {
		if status == rows[i].Status {
			fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%s\t%s\n", rows[i].Id, rows[i].Title, rows[i].Description, rows[i].Created, rows[i].Status)
		}
	}
	tabWriter.Flush()
}

func RemoveTask(id uint, rows *[]TableCl, verbose bool) {
	var nRows []TableCl
	for _, row := range *rows {
		if row.Id != id {
			if verbose {
				fmt.Printf("[FOUND ID] %d [TITLE] %s\n", row.Id, row.Title)
			}
			nRows = append(nRows, row)
		}
	}
	*rows = nRows
}

func IsValidStatus(s Status) bool {
	switch s {
	case StatusTodo:
		return true
	case StatusInProgress:
		return true
	case StatusDone:
		return true
	default:
		return false
	}
}
