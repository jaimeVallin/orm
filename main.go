package main

import (
	"net/http"

	// Importar los paquetes necesarios
	"inventario/go-backend/configs"
	"inventario/go-backend/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	// Inicializar la conexión a la base de datos
	configs.ConnectToDB()
}

func main() {
	// Crear una instancia del servidor Gin
	r := gin.Default()

	// Configurar las rutas relacionadas con las personas
	routes.PersonRouter(r)

	// Ruta de prueba en la raíz del servidor
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	// Iniciar el servidor en el puerto 8080 (por defecto)
	r.Run()
}
