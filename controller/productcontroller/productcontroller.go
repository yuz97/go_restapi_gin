package productcontroller

import (
	"net/http"
	"restapi_gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil ditampilkan", "product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil ditambahkan", "product": product})
}

func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "data tidak dapat diupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diupdate"})
}

func Destroy(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "data tidak dapat dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})

}
