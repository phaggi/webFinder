package services

import (
	"errors"
	"time"
	"webFinder/db"
	"webFinder/external"
	"webFinder/models"
)

type SearchService struct{}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (s *SearchService) TriggerScript(scriptName string) (int, error) {
	task := models.Task{
		ScriptName: scriptName,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	taskID, err := db.CreateTask(task)
	if err != nil {
		return 0, err
	}

	go func() {
		// Dummy execution of external program
		resultData := external.ExecutePIXRPA(scriptName)

		result := models.Result{
			TaskID:    taskID,
			Data:      resultData,
			CreatedAt: time.Now(),
		}

		db.SaveResults([]models.Result{result})
	}()

	return taskID, nil
}

func (s *SearchService) GetResults(taskID int) ([]models.Result, error) {
	results, err := db.GetResultsByTaskID(taskID)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("no results found")
	}

	return results, nil
}
