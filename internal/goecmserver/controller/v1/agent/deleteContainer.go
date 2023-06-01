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

func (a *AgentController) DeleteContainer(c *gin.Context) {
	var body []v1.ContainerOperation
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
		_ = json.NewEncoder(&buf).Encode(pending.ContainerId)
		url := fmt.Sprintf("http://%s/v1/container/delete", host)

		req, _ := http.NewRequest(http.MethodDelete, url, &buf)
		req.Close = true
		resp, _ := http.DefaultClient.Do(req)
		if resp != nil {
			resp.Body.Close()
		}
	}
	core.WriteResponse(c, nil, map[string]bool{"ack": true})
}
