package models

import "time"

type Task struct {
	ID         int       `json:"id"`
	ScriptName string    `json:"script_name"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
