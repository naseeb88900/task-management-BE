package controllers

import (
	"context"
	"fmt"
	"net/http"
	"task-management/models"
	"task-management/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Status = "To Do"

	collection := utils.GetCollection(utils.DB, "tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func GetTasks(c *gin.Context) {
	userID := c.GetString("userID")        
	userRole := c.GetString("userRole")   
	collection := utils.GetCollection(utils.DB, "tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var cursor *mongo.Cursor
	var err error

	fmt.Println("rol ------", userRole)
	if userRole == "manager" {
		cursor, err = collection.Find(ctx, bson.M{})
	} else {
		cursor, err = collection.Find(ctx, bson.M{"assignedto": userID})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}


func UpdateTask(c *gin.Context) {
	fmt.Println("Updating")
	taskID := c.Param("id")
	userID := c.GetString("userID")
	userRole := c.GetString("userRole")

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := utils.GetCollection(utils.DB, "tasks")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	fmt.Println("Object Id:", objID)
	fmt.Println("Assigned to:", userID)

	filter := bson.M{"_id": objID}

	if userRole != "manager" {
		filter["assignedto"] = userID
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":      updatedTask.Status,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	fmt.Println("Result", result)
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not assigned to you"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
    taskID := c.Param("id")
    userID := c.GetString("userID")
    userRole := c.GetString("userRole") 

    collection := utils.GetCollection(utils.DB, "tasks")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    filter := bson.M{"_id": objID}

    if userRole != "manager" {
        filter["assignedto"] = userID
    }

    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        return
    }

    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not assigned to you"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

