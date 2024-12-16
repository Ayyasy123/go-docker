package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User representasi data user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory data store
var users = []User{}
var nextID = 1

func main() {
	r := gin.Default()

	// Endpoint untuk mengambil seluruh user
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// Endpoint untuk mengambil seluruh user
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Endpoint untuk mengambil user berdasarkan ID
	r.GET("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
	})

	// Endpoint untuk menambahkan user baru
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = nextID
		nextID++
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	// Endpoint untuk mengupdate user berdasarkan ID
	r.PUT("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, user := range users {
			if user.ID == id {
				users[i].Name = updatedUser.Name
				users[i].Age = updatedUser.Age
				c.JSON(http.StatusOK, users[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
	})

	// Endpoint untuk menghapus user berdasarkan ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User dihapus"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
	})

	r.Run(":8080") // Jalankan server pada port 8080
}
