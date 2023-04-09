package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"hacktiv8_Project1_BookAPI/database"
	"hacktiv8_Project1_BookAPI/models"
)

func GetAllBook(ctx *gin.Context) {
	db := database.GetDB()
	var books []models.Book
	db.Find(&books)
	ctx.JSON(http.StatusOK, books)
}

func GetBookById(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("bookID")
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("Book with id %v not found", id),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("bookID")
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	err := db.Model(&book).Where("id = ?", id).Updates(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	db.Create(&book)
	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("bookID")
	var book models.Book

	if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})

}
