package main

import (
	"example.com/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run()
}
