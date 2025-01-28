package httpEngine

import (
	"example/web-service-gin/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

var l *logic.Logic

func Run(Port string) {
	router := gin.Default()

	var l = logic.Logic{}

	router.GET("/albums", l.GetAllAlbums)

	router.GET("/albums/:id", l.GetAlbumByID)

	router.POST("/albums", l.SaveAlbum)

	fmt.Println(router.Run("localhost:8080"))
}
