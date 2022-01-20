package controllers

import (
	"github.com/RyanCheungJF/CVWO-Backend/database"
	"github.com/RyanCheungJF/CVWO-Backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetTask(c *fiber.Ctx) error {
	id := c.Params("userid")
	var task []models.Task
	database.DB.Where("user_id = ?", id).Find(&task)
	return c.JSON(task)
}

func AddTask(c *fiber.Ctx) error {
	task := new(models.Task)
	err := c.BodyParser(task)
	if err != nil {
		c.Status(fiber.StatusServiceUnavailable)
		return c.JSON(fiber.Map{
			"message": "Failed to create task",
		})
	}
	database.DB.Create(&task)
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedTask := new(models.Task)
	err := c.BodyParser(updatedTask)
	if err != nil {
		c.Status(fiber.StatusServiceUnavailable)
		return c.JSON(fiber.Map{
			"message": "Failed to update task",
		})
	}

	var task models.Task
	database.DB.First(&task, id)
	if task.Name == "" {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "No task found with given ID",
		})
	}

	task.Status = updatedTask.Status
	database.DB.Save(&task)
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	database.DB.First(&task, id)
	if task.Name == "" {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "No task found with given ID",
		})
	}
	database.DB.Delete(&task)
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
