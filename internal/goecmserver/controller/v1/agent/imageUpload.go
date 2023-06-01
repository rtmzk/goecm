package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"strconv"
)

func (a *AgentController) ImageUpload(c *gin.Context) {
	fileName := c.PostForm("fileName")
	chunk, _ := strconv.Atoi(c.PostForm("chunk"))
	chunks, _ := strconv.Atoi(c.PostForm("chunks"))

	file, err := c.FormFile("file")
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "无法获取文件数据"), nil)
		return
	}

	err = c.SaveUploadedFile(file, "/tmp/"+fileName+strconv.Itoa(chunk))
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrImageUpload, "无法保存块数据"), nil)
		return
	}
	//src, err := file.Open()
	//if err != nil {
	//	core.WriteResponse(c, errors.WithCode(code.ErrUpload, "无法获取块数据"), nil)
	//	return
	//}
	//defer src.Close()
	//out, err := os.OpenFile("/tmp/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	core.WriteResponse(c, errors.WithCode(code.ErrUpload, "无法合并块数据"), nil)
	//	return
	//}
	//defer out.Close()
	//_, err = io.Copy(out, src)
	//if err != nil {
	//	core.WriteResponse(c, errors.WithCode(code.ErrUpload, "无法合并块数据到目标文件"), nil)
	//	return
	//}
	if chunk == chunks {
		core.WriteResponse(c, nil, map[string]string{"status": "end"})
	} else {
		core.WriteResponse(c, nil, map[string]string{"status": "uploading"})
	}

	return
}
