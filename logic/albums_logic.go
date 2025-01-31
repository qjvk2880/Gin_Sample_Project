package logic

import (
	"example/web-service-gin/model"
	"example/web-service-gin/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Logic struct{}

var albumRepo repository.Repository = repository.NewMemoryRepository()

func (l *Logic) GetAllAlbums(c *gin.Context) {
	albums, err := albumRepo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (l *Logic) GetAlbumByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 0, 64)
	album, err := albumRepo.GetByID(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (l *Logic) SaveAlbum(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	err := albumRepo.Save(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
