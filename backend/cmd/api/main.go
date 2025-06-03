package main

import (
	"firebase-playground/internal/interface/router"
)

func main() {
	router := router.NewRouter()
	router.Run(":8080")
}
