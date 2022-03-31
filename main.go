package main

import (
	config "boilerplate/configs"
	"boilerplate/routes"
	"os"
)

func main() {
	config.Init()
	config.InitRedis(1)
	r := routes.Init()
	r.Run(":" + os.Getenv("APP_PORT"))
}
