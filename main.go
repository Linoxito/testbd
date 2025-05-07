package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// Определяем модели
type Product struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Brand struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	db.AutoMigrate(&Product{}, &Brand{})

	r := gin.Default()

	// Маршруты для продуктов
	r.GET("/products", getProducts)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Маршруты для брендов
	r.GET("/brands", getBrands)
	r.POST("/brands", createBrand)
	r.PUT("/brands/:id", updateBrand)
	r.DELETE("/brands/:id", deleteBrand)

	r.Run(":8080")
}

// Функции работы с продуктами
func getProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func updateProduct(c *gin.Context) {
	var product Product
	id := c.Param("id")
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Продукт удален"})
}

// Функции работы с брендами
func getBrands(c *gin.Context) {
	var brands []Brand
	db.Find(&brands)
	c.JSON(http.StatusOK, brands)
}

func createBrand(c *gin.Context) {
	var brand Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&brand)
	c.JSON(http.StatusCreated, brand)
}

func updateBrand(c *gin.Context) {
	var brand Brand
	id := c.Param("id")
	if err := db.First(&brand, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Бренд не найден"})
		return
	}
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&brand)
	c.JSON(http.StatusOK, brand)
}

func deleteBrand(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&Brand{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Бренд удален"})
}
