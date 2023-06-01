package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/constand"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"net/http"
	"strings"
)

func (a *AgentController) DeleteImage(c *gin.Context) {
	var body []v1.ImageOperation
	if err := c.ShouldBindJSON(&body); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "客户端传参错误"), nil)
		return
	}

	for _, pending := range body {
		var buf bytes.Buffer
		var host = pending.Host
		for _, h := range constand.HeartbeatItem {
			if strings.Contains(h, host) {
				host = h
			}
		}
		_ = json.NewEncoder(&buf).Encode(pending.ImageId)
		url := fmt.Sprintf("http://%s/v1/image/delete", host)

		req, _ := http.NewRequest(http.MethodDelete, url, &buf)
		req.Close = true
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			core.WriteResponse(c, err, nil)
			return
		}
		if resp != nil {
			resp.Body.Close()
		}
	}
	core.WriteResponse(c, nil, map[string]bool{"ack": true})
}
