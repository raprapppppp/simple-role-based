package repository

import (
	"role-based/models"

	"gorm.io/gorm"
)


type TaskRepo interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTask(id int) ([]models.Task, error)
	DeleteTask(task models.Task) error 
	UpdateTask(task models.Task) (models.Task, error)

}

type TaskDbInjection struct{
	db *gorm.DB
}

//Task Repo initializer
func TaskRepoInit(db *gorm.DB) TaskRepo {
	return &TaskDbInjection{db}
}

func (r *TaskDbInjection) CreateTask(task models.Task) (models.Task, error) {

	err := r.db.Create(&task).Error

	if err != nil {
		return models.Task{}, err
	}
	
	return task, nil
}

func (r *TaskDbInjection) GetTask(id int) ([]models.Task, error){
	var task []models.Task

	err := r.db.Where("account_id = ?", id).Find(&task).Error
	if err != nil {
		return []models.Task{}, err
	}
	return task, nil
}

func (r *TaskDbInjection) DeleteTask(task models.Task) error {

	err := r.db.Delete(&task, task.ID).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *TaskDbInjection) UpdateTask(task models.Task) (models.Task, error){

	err := r.db.Where("id = ?", task.ID).Updates(&task).Error

	if err != nil {
		return models.Task{} ,err
	}
	return task, nil
}