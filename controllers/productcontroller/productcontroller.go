package productcontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"go-web-native/models/productmodel"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	products := productmodel.GetAll()

	c.HTML(http.StatusOK, "product.html", gin.H{
		"products": products,
	})
}

func Add(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/product/createprod.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Template parsing error: %v", err)
			return
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(c.Writer, data)
		return
	}

	if c.Request.Method == http.MethodPost {
		var product entities.Product

		categoryId, err := strconv.Atoi(c.PostForm("category_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category_id")
			return
		}

		stock, err := strconv.Atoi(c.PostForm("stock"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid stock")
			return
		}

		product.Name = c.PostForm("name")
		product.CategoryID = uint(categoryId)
		product.Stock = int(stock)
		product.Description = c.PostForm("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
			return
		}

		c.Redirect(http.StatusSeeOther, "/products")
	}
}

func Detail(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detailprod.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Template parsing error: %v", err)
		return
	}

	temp.Execute(c.Writer, data)
}

func Edit(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/product/editprod.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Template parsing error: %v", err)
			return
		}

		idString := c.Query("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()

		data := map[string]any{
			"product":    product,
			"categories": categories,
		}

		temp.Execute(c.Writer, data)
		return
	}

	if c.Request.Method == http.MethodPost {
		var product entities.Product

		idString := c.PostForm("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}

		categoryId, err := strconv.Atoi(c.PostForm("category_id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid category_id")
			return
		}

		stock, err := strconv.Atoi(c.PostForm("stock"))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid stock")
			return
		}

		product.Name = c.PostForm("name")
		product.CategoryID = uint(categoryId)
		product.Stock = int(stock)
		product.Description = c.PostForm("description")
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			c.Redirect(http.StatusTemporaryRedirect, c.Request.Referer())
			return
		}

		c.Redirect(http.StatusSeeOther, "/products")
	}
}

func Delete(c *gin.Context) {
	idString := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := productmodel.Delete(id); err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete product: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/products")
}
