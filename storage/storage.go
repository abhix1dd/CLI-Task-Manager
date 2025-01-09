package storage

import (
    "cli-task-manager/models"
    "encoding/json"
    // "io/ioutil"
    "os"
)

func LoadTasks(filePath string) ([]models.Task, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return []models.Task{}, nil // Return an empty list if the file doesn't exist
    }
    defer file.Close()

    var tasks []models.Task
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    return tasks, err
}

func SaveTasks(filePath string, tasks []models.Task) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(tasks)
}
