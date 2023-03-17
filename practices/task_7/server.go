package main

import (
	"Product/routers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routers.Api()
}
