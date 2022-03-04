package middleware

import (
	"cfttest/workwithfiles"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetFileList(c *gin.Context) {
	c.JSON(200, workwithfiles.ListDirectory())
}
func GetFile(c *gin.Context) {
	name := c.Param("name")
	if exists, _ := workwithfiles.FileExists("tmp/" + name); exists {
		c.File("tmp/" + name)
	} else {
		c.String(http.StatusBadRequest, fmt.Sprintf("'%s' doesn't exist!", name))
	}
}
func AddFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	name := file.Filename
	if exists, _ := workwithfiles.FileExists("tmp/" + name); exists {
		c.String(http.StatusBadRequest, fmt.Sprintf("'%s' already exists!", name))
		return
	}
	returnedStatus := http.StatusOK
	var msg string
	err = c.SaveUploadedFile(file, "tmp/"+file.Filename)
	if err != nil {
		returnedStatus = http.StatusBadRequest
		msg = fmt.Sprintf("Something went wrong with %v", file.Filename)
	}
	c.String(returnedStatus, msg)
}
func UpdateFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	realFile, _ := file.Open()
	defer realFile.Close()
	bytes, _ := ioutil.ReadAll(realFile)
	newHash := fmt.Sprintf("%x", md5.Sum(bytes))
	oldHash, hashErr := workwithfiles.CalculateHash(file.Filename)
	if errors.Is(hashErr, os.ErrNotExist) {
		c.String(http.StatusBadRequest, "No such file")
		return
	}
	if newHash == oldHash {
		c.String(http.StatusOK, "Files are equal, no update required")
		return
	}
	if err = os.Remove("tmp/" + file.Filename); err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
		return
	}
	err = c.SaveUploadedFile(file, "tmp/"+file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
		return
	}
	c.String(http.StatusOK, "")

}
func DeleteFile(c *gin.Context) {
	name := c.Param("name")
	if exists, _ := workwithfiles.FileExists("tmp/" + name); exists {
		os.Remove("tmp/" + name)
		c.String(200, "")
	} else {
		c.String(http.StatusBadRequest, "No such file")
	}

}
