package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataRemover interface {
	DeleteAll()
}

func Delete(remover DataRemover) gin.HandlerFunc {
	return func(context *gin.Context) {

		remover.DeleteAll()
		m := map[string]string{
			"status": "OK",
		}

		context.JSON(http.StatusOK, m)

		return
	}
}
