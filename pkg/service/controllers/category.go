package controllers

import (
	"app/database"
	"net/http"

	"app/pkg/service/repositories"
	"app/pkg/service/responses"

	"github.com/gin-gonic/gin"
)

func FindCategoriesByName(c *gin.Context) {
	name := c.Param("name")

	db, error := database.Connection()
	if error != nil {
		responses.Error(c, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repo := repositories.NewCategoryRepository(db)
	categories, error := repo.FindCategoryByName(name)
	if error != nil {
		responses.Error(c, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(c, http.StatusOK, categories)

}
