package books

import (
	"net/http"

	"github.com/isravm89w/gin-rest-api-example/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, &books)
}
