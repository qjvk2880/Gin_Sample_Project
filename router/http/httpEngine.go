package httpEngine

import (
	"example/web-service-gin/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

var l *logic.Logic

func Run(Port string) {
	router := gin.Default()

	var albumLogic = logic.Logic{}

	router.GET("/albums", albumLogic.GetAllAlbums)

	router.GET("/albums/:id", albumLogic.GetAlbumByID)

	router.POST("/albums", albumLogic.SaveAlbum)

	fmt.Println(router.Run("localhost:8080"))
}
