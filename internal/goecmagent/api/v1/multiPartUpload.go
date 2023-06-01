package v1

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ecm/utils"
	"io/ioutil"
	"os"
)

func MultiPartUpload(c *gin.Context) {
	fileName := c.Query("name")
	path := c.Query("path")
	chunk := c.Query("chunk")

	if utils.IsNotExist(path) {
		os.MkdirAll(path, 0755)
	}

	if chunk == "0" && utils.IsExist(path+fileName) {
		_ = os.RemoveAll(path + fileName)
	}

	targetFile, err := os.OpenFile(path+fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	targetWriter := bufio.NewWriter(targetFile)

	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = targetWriter.Write(buf)
	_ = targetWriter.Flush()

	targetFile.Close()
	return
}
