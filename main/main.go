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
	logic.InitRedis()
	http.Run(":8081")
}
