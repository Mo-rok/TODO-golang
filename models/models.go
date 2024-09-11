package models

import "time"

type List struct {
	Name string
}
type Task struct {
	Name        string
	Description string
	StartAt     time.Time
	EndAt       time.Time
	List        *List
}
