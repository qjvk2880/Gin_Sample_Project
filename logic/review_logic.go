package logic

import (
	"example/web-service-gin/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReviewLogic struct{}

var reviewRepo repository.ReviewRepository = repository.NewReviewRdbRepository()

func (l *ReviewLogic) GetAllReviewsByAlbumId(c *gin.Context) {
	albumId, _ := strconv.ParseUint(c.Param("id"), 0, 64)

	reviews, err := reviewRepo.GetAllByAlbum(uint(albumId))

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "review not found"})
	}

	c.IndentedJSON(http.StatusOK, reviews)
}
