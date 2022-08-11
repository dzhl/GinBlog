package main

import (
	"ginblog/model"
	"ginblog/routes"
	//"golang.org/x/crypto/bcrypt"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
