package main

import (
	"inventario/go-backend/configs"
	"inventario/go-backend/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}
