package main

import (
	"gin_sample/router"
)

func main() {
	router := router.GetRouter()
	router.Run(":3000")
}
