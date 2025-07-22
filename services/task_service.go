package services

import (
	"role-based/models"
	"role-based/repository"
)

type TaskServices interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTask(id int) ([]models.Task, error)
	DeleteTask(task models.Task) (string, error)
	UpdateTask(task models.Task) (models.Task, error)
}

type TaskRepoInjection struct {
	repo repository.TaskRepo
}

func TaskServicesInit(repo repository.TaskRepo) TaskServices {
	return &TaskRepoInjection{repo}
}

func(s *TaskRepoInjection) CreateTask(task models.Task) (models.Task, error) {
	return s.repo.CreateTask(task)
}	

func(s *TaskRepoInjection) GetTask(id int) ([]models.Task, error){
	return s.repo.GetTask(id)
}

func (s *TaskRepoInjection) DeleteTask(task models.Task) (string, error){
	
	err := s.repo.DeleteTask(task)
	if err != nil {
		return "", err
	}
	return "Deleted", nil
}

func(s *TaskRepoInjection) UpdateTask(task models.Task) (models.Task, error){
	return s.repo.UpdateTask(task)
}
