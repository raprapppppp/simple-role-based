package handlers

import (
	"role-based/models"
	"role-based/services"

	"github.com/gofiber/fiber/v2"
)

type TaskHandlerInjection struct{
	service services.TaskServices
}

func TaskHandlerInit (service services.TaskServices) *TaskHandlerInjection{
	return &TaskHandlerInjection{service}
}

//Create Task
func (h *TaskHandlerInjection) CreateTask(t *fiber.Ctx) error{
	var task models.Task

	err := t.BodyParser(&task)
	if err != nil {
		return t.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"}) 
	}

	userId := t.Locals("id")
	if userId == nil {
		return t.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	idFloat, ok := userId.(float64)
	if !ok {
		return t.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user ID type"})
	}

	task.AccountId = uint(idFloat)

	tasks, err := h.service.CreateTask(task)
	if err !=nil {
		return t.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
	}

	return t.Status(fiber.StatusOK).JSON(tasks)
}

//Get Task base on who is the user
func(h *TaskHandlerInjection) GetTask(t *fiber.Ctx) error {

	userId := t.Locals("id")
	if userId == nil {
		return t.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	idFloat, ok := userId.(float64)
	if !ok {
		return t.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid user ID type"})
	}

	tasks, err := h.service.GetTask(int(idFloat))
	if err != nil {
		return err
	}

	return t.Status(fiber.StatusOK).JSON(tasks)
}


//Delete
func (h *TaskHandlerInjection) DeleteTask(t *fiber.Ctx) error {
	var task models.Task

	err := t.BodyParser(&task)
	if err != nil{
		return t.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"}) 
	}

	mess,err := h.service.DeleteTask(task)

	if err != nil {
		return err
	}

	return t.Status(fiber.StatusOK).JSON(fiber.Map{"error": mess}) 
}

//Update
func (h *TaskHandlerInjection) UpdateTask(t *fiber.Ctx) error {
	var task models.Task

	err := t.BodyParser(&task)
	if err != nil {
		return t.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"}) 
	}

	upTask, err := h.service.UpdateTask(task)
	if err != nil {
		return err
	}
	return t.Status(fiber.StatusOK).JSON(upTask)

}