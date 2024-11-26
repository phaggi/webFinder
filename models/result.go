package models

import "time"

type Result struct {
	ID        int       `json:"id"`
	TaskID    int       `json:"task_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}
