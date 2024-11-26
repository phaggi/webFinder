package db

import (
	"database/sql"
	"fmt"
	"log"
	"webFinder/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}

func CreateTask(task models.Task) (int, error) {
	var taskID int
	err := db.QueryRow("INSERT INTO tasks (script_name, status, created_at) VALUES ($1, $2, $3) RETURNING id",
		task.ScriptName, task.Status, task.CreatedAt).Scan(&taskID)
	if err != nil {
		return 0, err
	}

	return taskID, nil
}

func SaveResults(results []models.Result) error {
	for _, result := range results {
		_, err := db.Exec("INSERT INTO results (task_id, data, created_at) VALUES ($1, $2, $3)",
			result.TaskID, result.Data, result.CreatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetResultsByTaskID(taskID int) ([]models.Result, error) {
	rows, err := db.Query("SELECT id, task_id, data, created_at FROM results WHERE task_id = $1", taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var result models.Result
		err := rows.Scan(&result.ID, &result.TaskID, &result.Data, &result.CreatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
