package agent

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmagent/static"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/resource"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (a *AgentController) EnvCheckPrepare(c *gin.Context) {
	var buf v1.CheckRules

	err := json.Unmarshal(static.CHECK_RULES, &buf)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, &buf)
}

func (a *AgentController) EnvCheckAction(c *gin.Context) {
	var agents []v1.Agent
	var result v1.CheckRules
	var buf v1.CheckRules
	var ips []string

	if err := c.ShouldBindJSON(&agents); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "传参错误"), nil)
		return
	}

	_ = json.Unmarshal(resource.CHECK_RULES, &result)
	_ = json.Unmarshal(resource.CHECK_RULES, &buf)

	for _, agent := range agents {
		ips = append(ips, agent.AgentAddr+":"+strconv.Itoa(agent.AgentPort))
	}

	for _, endpoint := range ips {
		var temp = new(v1.CheckRules)
		url := fmt.Sprintf("http://%s/v1/envc", endpoint)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Close = true
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrNetConnection, "无法获取检查信息"), nil)
			break
		}
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, &temp.Rules)

		for idx, item := range temp.Rules {
			if buf.Rules[idx].Status == "none" {
				buf.Rules[idx].Status = "OK"
				buf.Rules[idx].Message = ""
			}
			if item.Status == "FAILED" {
				buf.Rules[idx].Status = item.Status
				if buf.Rules[idx].Message == "" {
					buf.Rules[idx].Message = item.Message
				} else {
					buf.Rules[idx].Message = fmt.Sprintf("%s  %s", buf.Rules[idx].Message, item.Message)
				}
			} else if buf.Rules[idx].Status == "FAILED" {
				continue
			}
		}
	}

	core.WriteResponse(c, nil, buf.Rules)
}
