package httpEngine

import (
	"example/web-service-gin/logic"
	"example/web-service-gin/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var l *logic.Logic

func Run(Port string) {
	router := gin.Default()

	repo := repository.NewMemoryRepository()

	l = logic.NewLogic(repo)

	router.GET("/albums", func(c *gin.Context) {
		albums, err := l.GetAllAlbums()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		}

		c.IndentedJSON(http.StatusOK, albums)
	})

	router.GET("/albums/:id", func(c *gin.Context) {
		id := c.Param("id")
		album, err := l.GetAlbumByID(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		}

		c.IndentedJSON(http.StatusOK, album)
	})

	router.POST("/albums", func(c *gin.Context) {
		var newAlbum repository.Album
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		err := l.SaveAlbum(newAlbum)
		if err != nil {
			return
		}

		c.IndentedJSON(http.StatusCreated, newAlbum)
	})

	fmt.Println(router.Run("localhost:8080"))
}
