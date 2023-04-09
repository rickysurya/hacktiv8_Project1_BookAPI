package routers

import (
	"github.com/gin-gonic/gin"
	"hacktiv8_Project1_BookAPI/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetAllBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.GET("/books/:bookID", controllers.GetBookById)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
