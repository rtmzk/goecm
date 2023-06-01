package agent

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"net/http"
)

func (a *AgentController) PullImage(c *gin.Context) {
	var data v1.ImagePull
	var err error
	var resp *http.Response
	if err = c.ShouldBindJSON(&data); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "传参错误"), nil)
		return
	}
	imageName := data.ImageName

	url := fmt.Sprintf("http://%s/v1/image/pull?name=%s", data.Node, imageName)
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Close = true
	resp, err = http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp != nil {
			resp.Body.Close()
		}
		core.WriteResponse(c, errors.WithCode(code.ErrImagePullWithConnectionErrWrapOK, "拉取镜像失败"), nil)
		return
	}
	c.DataFromReader(http.StatusOK, resp.ContentLength, "application/json", resp.Body, nil)
}
