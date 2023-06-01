package docker

import (
	"fmt"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
)

func (d *DockerController) ImageUpload(c *gin.Context) {
	// log.Infof("%+v", c.Request.Header)
	log.L(c).Debug("image upload function called.")
	var saveName string
	var fs []string
	var savePath = "/data"

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["file"]

	for _, file := range files {
		if utils.IsEmpty(file.Filename) {
			core.WriteResponse(c, errors.WithCode(code.ErrBind, "请选择文件"), nil)
			return
		}

		filename := filepath.Base(file.Filename)
		randname := utils.RandString(utils.Alphabet36, 15)
		saveName = randname + path.Ext(filename)
		savePath = savePath + saveName
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrSystemStorageIssue, "无法保存文件"), nil)
		}
		log.Infof("Recive file: %s, size: %d", file.Filename, file.Size)
		fs = append(fs, savePath)
	}

	err = d.srv.Docker().ImageLoad(c, fs)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
