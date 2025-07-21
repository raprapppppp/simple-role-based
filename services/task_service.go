package services

import (
	"role-based/models"
	"role-based/repository"
)

type TaskServices interface {
	CreateTask(task models.Task) error
	GetTask(id int) ([]models.Task, error)
}

type TaskRepoInjection struct {
	repo repository.TaskRepo
}

func TaskServicesInit(repo repository.TaskRepo) TaskServices {
	return &TaskRepoInjection{repo}
}

func(s *TaskRepoInjection) CreateTask(task models.Task) error {
	return s.repo.CreateTask(task)
}

func(s *TaskRepoInjection) GetTask(id int) ([]models.Task, error){
	return s.repo.GetTask(id)
}