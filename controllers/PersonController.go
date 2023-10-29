package controllers

import (
	"inventario/go-backend/configs"
	"inventario/go-backend/models"

	"github.com/gin-gonic/gin"
)

// Estructura que representa el cuerpo de la solicitud JSON para crear una persona.
type PersonRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   uint   `json:"phone"`
}

// PersonCreate maneja la creación de una nueva persona.
func PersonCreate(c *gin.Context) {
	// Parsear el cuerpo JSON de la solicitud en la estructura PersonRequestBody.
	body := PersonRequestBody{}
	c.BindJSON(&body)

	// Crear un nuevo objeto Person a partir de los datos recibidos.
	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	// Insertar el objeto Person en la base de datos.
	result := configs.DB.Create(&person)

	// Verificar si hubo errores en la inserción.
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &person)
}

// PersonGet recupera todas las personas de la base de datos.
func PersonGet(c *gin.Context) {
	var persons []models.Person
	configs.DB.Find(&persons)
	c.JSON(200, &persons)
	return
}

// PersonGetById recupera una persona por su ID.
func PersonGetById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)
	c.JSON(200, &person)
	return
}

// PersonUpdate actualiza una persona existente por su ID.
func PersonUpdate(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)

	// Parsear el cuerpo JSON de la solicitud en la estructura PersonRequestBody.
	body := PersonRequestBody{}
	c.BindJSON(&body)

	// Crear un objeto Person actualizado con los datos del cuerpo JSON.
	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	// Realizar la actualización en la base de datos.
	result := configs.DB.Model(&person).Updates(data)

	// Verificar si hubo errores en la actualización.
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &person)
}

// PersonDelete elimina una persona por su ID.
func PersonDelete(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.Delete(&person, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
