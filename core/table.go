package core

import (
	"time"
)

var (
	lastId uint = 0
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type tableCl struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Status      Status    `json:"status"`
}

func NewRow(title string, desc string, status Status) *tableCl {
	lastId++
	return &tableCl{
		Id:          lastId,
		Title:       title,
		Description: desc,
		Created:     time.Now(),
		Status:      status,
	}
}

func isValidStatus(s Status) bool {
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
