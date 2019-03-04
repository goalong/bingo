package main

import (
	"github.com/goalong/bingo/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8000")
}
