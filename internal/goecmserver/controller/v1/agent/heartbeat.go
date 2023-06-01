package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/constand"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"strconv"
	"strings"
)

func (a *AgentController) Heartbeat(c *gin.Context) {
	log.L(c).Debugf("heart beat function called. Host: %s", c.Request.Host)
	var heartbeatItem v1.Heartbeat
	var flag bool
	if err := c.ShouldBindJSON(&heartbeatItem); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "传参错误"), nil)
		return
	}

	for _, value := range constand.HeartbeatItem {
		if strings.Contains(value, heartbeatItem.AgentAddr) {
			flag = true
		}
	}

	if !flag {
		constand.HeartbeatItem = append(constand.HeartbeatItem, heartbeatItem.AgentAddr+":"+strconv.Itoa(heartbeatItem.AgentPort))
	}
	core.WriteResponse(c, nil, map[string]bool{"ack": true})
}
