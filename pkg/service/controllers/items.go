package controllers

import (
	"app/database"
	"net/http"
	"strconv"

	"app/pkg/service/repositories"
	"app/pkg/service/responses"

	"github.com/gin-gonic/gin"
)

func FindItemsByName(c *gin.Context) {
	name := c.Param("name")

	db, error := database.Connection()
	if error != nil {
		responses.Error(c, http.StatusInternalServerError, error)
		return
	}

	repo := repositories.NewItemRepository(db)
	items, error := repo.FindItemByName(name)
	if error != nil {
		responses.Error(c, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	responses.JSON(c, http.StatusOK, items)
}

func FindItemsByCategory(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("category"))

	db, error := database.Connection()
	if error != nil {
		responses.Error(c, http.StatusInternalServerError, error)
		return
	}

	repo := repositories.NewItemRepository(db)
	items, error := repo.FindItemByCategory(categoryID)
	if error != nil {
		responses.Error(c, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	responses.JSON(c, http.StatusOK, items)
}
