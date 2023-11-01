package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var (
	// postgresDBInstance is an instance for of postgres db in gorm format
	postgresDBInstance *gorm.DB
)

type Student struct {
	Id        int
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

func main() {
	// Database
	db := connectDB("0.0.0.0", "5432", "postgres", "mypass", "gin_gorm")

	// gin server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// POST
	r.POST("/students", func(c *gin.Context) {
		var createStudent Student
		c.Bind(&createStudent)

		student := Student{
			FirstName: createStudent.FirstName,
			LastName:  createStudent.LastName,
			Age:       createStudent.Age,
			Sex:       createStudent.Sex,
			Email:     createStudent.Email,
		}
		db.Create(&student)

		c.JSON(200, student)
		return
	})

	// GET
	r.GET("/students", func(c *gin.Context) {

	})
	//PUT
	//DELETE
	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func connectDB(host, port, user, password, dbName string) *gorm.DB {
	var err error
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)
	if postgresDBInstance == nil {
		postgresDBInstance, err = gorm.Open("postgres", connection)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("database is connected")
	}
	return postgresDBInstance
}
