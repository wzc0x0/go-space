package routers

import (
	"github.com/gin-gonic/gin"

	"go-space/constants"
	"go-space/models"
	"go-space/servers"
)

func GetMovies(c *gin.Context) {
	var moviesDto models.PaginationDto
	if err := c.ShouldBind(&moviesDto); err != nil {
		servers.ResponseException(c, constants.InvalidParams, err.Error())
		return
	}

	var total int
	var movies []models.SpMovie

	db := servers.DB.Model(&movies)
	db.Count(&total)
	offset := (moviesDto.Current - 1) * moviesDto.PageSize

	if offset >= total {
		servers.ResponseException(c, 1002, "end")
		return
	}

	if db.Offset(offset).Limit(moviesDto.PageSize).Find(&movies).Error != nil {
		servers.ResponseException(c, constants.ERROR, "")
		return
	}
	servers.ResponseOK(c, &models.Pagination{
		PaginationDto: moviesDto,
		Total:         total,
		List:          movies,
	})
}
