package router

import (
	"cfttest/middleware"

	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {

	r := gin.Default()
	r.GET("/getFileList", middleware.GetFileList)
	r.GET("/getFile/:name", middleware.GetFile)
	r.PUT("/putFile", middleware.AddFile)
	r.POST("/updateFile", middleware.UpdateFile)
	r.DELETE("/deleteFile/:name", middleware.DeleteFile)

	return r
}
