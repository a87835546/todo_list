package main

import (
	"log"
	"todo_list/internal/logic"
	"todo_list/internal/router"
)

func main() {
	log.Printf("main starting.....")

	http := router.InitRouter()
	logic.InitDb()
	err := logic.InitRedis()
	if err != nil {
		return
	}
	logic.InitMongoDB()
	err = http.Run(":8081")
	if err != nil {
		return
	}
}
