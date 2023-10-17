package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-server/internal/dto"
)

type DataGetter interface {
	GetAll() []dto.Request
}

func GetData(getter DataGetter) gin.HandlerFunc {
	return func(context *gin.Context) {
		data := getter.GetAll()
		context.JSON(http.StatusOK, data)
	}
}
