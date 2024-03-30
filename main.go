package main

import (
	"cicd_api/commons/router"
)

func main() {
	router := router.Router()
	router.Run()
}
