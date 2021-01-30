package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Riki-Okunishi/todo-api/src/model"
)

// タスク一覧
func TasksGET(c *gin.Context){
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// json返却用
	tasks := []model.Task{}
	for result.Next() {
		task := model.Task{}
		var id uint
		var createdAt, updatedAt time.Time
		var title, content string

		err = result.Scan(&id, &createdAt, &updatedAt, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		task.Content = content

		tasks= append(tasks, task)
	}
	
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// タスク検索
func FindByID(id uint) model.Task {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	task := model.Task{}
	for result.Next() {
		var createdAt, updatedAt time.Time
		var title, content string

		err = result.Scan(&id, &createdAt, &updatedAt, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		task.Content = content
	}
	return task
}

// タスク登録 POST(application/json)
func TaskPOST(c *gin.Context) {
	db := model.DBConnect()

	json := model.Task{}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()

	_, err := db.Exec("INSERT INTO task (title, content, created_at, updated_at) VALUES(?, ?, ?, ?)", json.Title, json.Content, now, now)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("POST sent.\ntitle: %s, content: %s\n", json.Title, json.Content)
}

// タスク更新
func TaskPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	content := c.PostForm("content")
	now := time.Now()

	_, err := db.Exec("UPDATE task SET title = ?, content = ?, updated_at = ? WHERE id = ?", title, content, now, id)
	if err != nil {
		panic(err.Error())
	}

	task := FindByID(uint(id))

	fmt.Println(task)
	c.JSON(http.StatusOK, gin.H{"task:": task})
}

// タスク削除
func TaskDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	_, err := db.Query("DELETE FROM task WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, "deleted")
}