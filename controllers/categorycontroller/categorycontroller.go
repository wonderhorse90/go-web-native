package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	categories := categorymodel.GetAll()

	// Render the category index page with the categories data
	c.HTML(http.StatusOK, "category.html", gin.H{
		"categories": categories,
	})
}

func Add(c *gin.Context) {
	if c.Request.Method == "GET" {
		// Render the create category page
		c.HTML(http.StatusOK, "views/category/createcat.html", nil)
	} else if c.Request.Method == "POST" {
		var category entities.Category

		// Get form values
		category.Name = c.DefaultPostForm("name", "")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		// Create category
		ok := categorymodel.Create(category)
		if !ok {
			// Render the create category page again in case of error
			c.HTML(http.StatusOK, "views/category/createcat.html", nil)
			return
		}

		// Redirect to the category list
		c.Redirect(http.StatusSeeOther, "/categories")
	}
}

func Edit(c *gin.Context) {
	if c.Request.Method == "GET" {
		// Get category ID from query
		idString := c.DefaultQuery("id", "")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category ID")
			return
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}

		// Render the edit category page with category data
		c.HTML(http.StatusOK, "views/category/editcat.html", data)
	} else if c.Request.Method == "POST" {
		var category entities.Category

		// Get category ID from form
		idString := c.DefaultPostForm("id", "")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category ID")
			return
		}

		// Get form values
		category.Name = c.DefaultPostForm("name", "")
		category.UpdatedAt = time.Now()

		// Update category
		if ok := categorymodel.Update(id, category); !ok {
			c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
			return
		}

		// Redirect to the category list
		c.Redirect(http.StatusSeeOther, "/categories")
	}
}

func Delete(c *gin.Context) {
	idString := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid category ID")
		return
	}

	// Delete the category
	if err := categorymodel.Delete(id); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting category: %v", err)
		return
	}

	// Redirect to the category list
	c.Redirect(http.StatusSeeOther, "/categories")
}
