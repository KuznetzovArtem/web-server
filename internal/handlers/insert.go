package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"web-server/internal/dto"
)

type DataSetter interface {
	Add(data dto.Request)
}

func InsertData(getter DataSetter) gin.HandlerFunc {
	return func(context *gin.Context) {
		body, err := io.ReadAll(context.Request.Body)
		defer context.Request.Body.Close()
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		var data dto.Request
		if err := json.Unmarshal(body, &data); err != nil {
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		getter.Add(data)
		ok := map[string]string{
			"status": "OK",
		}

		context.JSON(http.StatusOK, ok)
	}
}
