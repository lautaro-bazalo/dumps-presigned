package main

import (
	"dumps-presigned/api"
	"dumps-presigned/internal/application"
	"dumps-presigned/internal/presigner"
	"github.com/gin-gonic/gin"
	"os"
)

type dumps struct {
	presigner  *presigner.Presigner
	bucketName string
}

func main() {

	app := application.NewApplication()
	bucketName := os.Getenv("bucket-name")

	dumps := dumps{
		presigner:  app.Presigner,
		bucketName: bucketName,
	}

	r := gin.Default()
	r.GET("/dumps", dumps.getDumps)

	if err := r.Run(":9290"); err != nil {
		panic(err)
	}

}

func (d dumps) getDumps(ctx *gin.Context) {

	req := api.DumpRequest{}

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	reqPresign, err := d.presigner.GetObject(d.bucketName, req.Path, req.Timeout)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"url": reqPresign.URL,
	})

}
