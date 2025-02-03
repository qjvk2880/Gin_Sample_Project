package httpEngine

import (
	"example/web-service-gin/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

var l *logic.AlbumLogic

func Run(Port string) {
	router := gin.Default()

	var albumLogic = logic.AlbumLogic{}
	var reviewLogic = logic.ReviewLogic{}

	router.GET("/albums", albumLogic.GetAllAlbums)

	router.GET("/albums/:id", albumLogic.GetAlbumByID)

	router.POST("/albums", albumLogic.SaveAlbum)

	router.GET("/reviews/:id", reviewLogic.GetAllReviewsByAlbumId)

	fmt.Println(router.Run("localhost:8080"))
}
