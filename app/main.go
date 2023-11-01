package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

var (
	// postgresDBInstance is an instance for of postgres db in gorm format
	postgresDBInstance *gorm.DB
)

func main() {

	// Database
	//db := connectDB("postgres", "5432", "postgres", "mypass", "gin_gorm")

	// gin server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/students", func(context *gin.Context) {

	})
	r.GET("/students", func(context *gin.Context) {

	})
	//PUT

	//DELETE
	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func create() {}
func update() {}
func delete() {}
func get()    {}

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
			panic(err)
		}
		log.Println("database is connected")
	}
	return postgresDBInstance
}
