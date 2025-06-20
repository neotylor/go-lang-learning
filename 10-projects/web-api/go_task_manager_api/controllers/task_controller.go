package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/database"
	"github.com/neotylor/go-lang-learning/tree/master/10-projects/web-api/go_task_manager_api/models"

	"github.com/gorilla/mux"
)

// GetTasks godoc
// @Summary      List all tasks
// @Description  Get all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

// GetTask godoc
// @Summary      Get a task by ID
// @Description  Get a task by its ID
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  models.Task
// @Failure      404  {string}  string  "Task not found"
// @Router       /tasks/{id} [get]
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Create a new task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      models.Task  true  "Task Body"
// @Success      201   {object}  models.Task
// @Router       /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Create(&task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update a task by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Task ID"
// @Param        task  body      models.Task true  "Task Body"
// @Success      200   {object}  models.Task
// @Failure      404   {string}  string  "Task not found"
// @Router       /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Delete a task by its ID
// @Tags         tasks
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {string}  string  "Task not found"
// @Router       /tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task models.Task
	result := database.DB.Delete(&task, id)
	if result.RowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}
