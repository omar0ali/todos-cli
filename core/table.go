// Package core
package core

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/xeonx/timeago"
)

var LastID uint = 0

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type TableCl struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Status      Status    `json:"status"`
}

func (t *TableCl) GetTime() string {
	return timeago.English.Format(t.Created)
}

func NewRow(title string, desc string, status Status) *TableCl {
	LastID++
	return &TableCl{
		ID:          LastID,
		Title:       title,
		Description: desc,
		Created:     time.Now(),
		Status:      status,
	}
}

func DisplayRecords(rows []TableCl, verbose bool) {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	tabWriter.Write([]byte("ID\tTitle\tDscription\tCreated\tStatus\n"))
	for i := range rows {
		fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%s\t%s\n", rows[i].ID, rows[i].Title, rows[i].Description, rows[i].GetTime(), rows[i].Status)
	}
	tabWriter.Flush()
}

func DisplayRecordsStatus(rows []TableCl, verbose bool, status Status) {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	tabWriter.Write([]byte("ID\tTitle\tDescription\tCreated\tStatus\n"))
	for i := range rows {
		if status == rows[i].Status {
			fmt.Fprintf(tabWriter, "%v\t%s\t%s\t%s\t%s\n", rows[i].ID, rows[i].Title, rows[i].Description, rows[i].GetTime(), rows[i].Status)
		}
	}
	tabWriter.Flush()
}

func RemoveTask(id uint, rows *[]TableCl, verbose bool) {
	var nRows []TableCl
	for _, row := range *rows {
		if row.ID != id {
			if verbose {
				log.Printf("[FOUND ID] %d [TITLE] %s\n", row.ID, row.Title)
			}
			nRows = append(nRows, row)
		}
	}
	*rows = nRows
}

func (t *TableCl) updateStatus(status Status) {
	t.Status = status
}

func UpdateStatusTask(id uint, status Status, rows []TableCl, verbose bool) {
	for i := range rows {
		if (rows)[i].ID == id {
			if verbose {
				log.Printf("[FOUND ID] %d [TITLE] %s\n", rows[i].ID, rows[i].Title)
			}
			rows[i].updateStatus(status)
			break
		}
	}
}

func EditTask(rows []TableCl, id uint, title string, desc string, verbose bool) {
	for i := range rows {
		if rows[i].ID == id {
			if verbose {
				log.Printf("[FOUND ID] %d [OLD TITLE] %s [OLD DESCRIPTION] %s\n", rows[i].ID, rows[i].Title, rows[i].Description)
			}
			if title != "" {
				rows[i].editTitle(title)
			}
			if desc != "" {
				rows[i].editDesc(desc)
			}
			break
		}
	}
}

func (t *TableCl) editTitle(title string) {
	t.Title = title
}

func (t *TableCl) editDesc(desc string) {
	t.Description = desc
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
