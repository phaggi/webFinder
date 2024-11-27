package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"webFinder/models"

	_ "github.com/lib/pq"
)

type Config struct {
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
}

var db *sql.DB

func InitDB() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}

func loadConfig() (*Config, error) {
	configPath := ".secrets/config.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", configPath)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
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
