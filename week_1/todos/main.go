package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	// "syreclabs.com/go/faker"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

var db *gorm.DB

func main() {
	// fmt.Println("vim-go")
	var err error
	// id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:8306)/todos?parseTime=true")

	if err != nil {
		log.Fatal("Thien Failed to connect DB")
	}

	router := gin.Default()

	router.GET("/todos", listTodos)
	router.POST("/todos", createTodo)
	router.Run(":8088")

	db.LogMode(true)

	// defer db.Close()

	// err = db.AutoMigrate(Todo{}).Error

}

func createTodo(c *gin.Context) {
	var argument struct {
		Title string
	}

	err := c.BindJSON(&argument)

	if err != nil {
		c.String(400, "invalid params")
		return
	}

	todo := Todo{
		Title: argument.Title,
		// Title: faker.Address().State(),
		// CreatedAt: time.Now(),
	}

	err = db.Create(&todo).Error // should be & to take id

	if err != nil {
		c.String(500, "failed to create new todo", err)
		return
	}

	c.JSON(200, todo)
}

func listTodos(c *gin.Context) {
	var todos []Todo

	err := db.Find(&todos).Error

	if err != nil {
		c.String(500, "failed to list todoList")
		return
	}

	c.JSON(200, todos)
}
