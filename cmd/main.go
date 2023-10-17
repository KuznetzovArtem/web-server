package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"
	"web-server/internal/dto"
	"web-server/internal/handlers"
	"web-server/internal/repository"
)

func main() {
	data := make(map[string]string, 0)

	values := os.Environ()
	for _, v := range values {
		splited := strings.Split(v, "=")
		data[splited[0]] = splited[1]
	}

	dsn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		data["USER"],
		data["PASSWORD"],
		data["HOST"],
		data["PORT"],
		data["DB_NAME"],
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.AutoMigrate(&dto.Request{})
	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	repo := repository.NewRequestWriterRepo(db)

	r.POST("/insert", handlers.InsertData(repo))
	r.GET("/get", handlers.GetData(repo))
	r.DELETE("/delete", handlers.Delete(repo))

	err = r.Run(":" + data["APP_PORT"])
	if err != nil {
		fmt.Println(err)
		return
	}
}
