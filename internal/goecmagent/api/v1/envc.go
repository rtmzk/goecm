package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-ecm/internal/goecmagent/model/v1"
	"go-ecm/internal/goecmagent/static"
	"go-ecm/pkg/core"
	"go-ecm/utils"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Envc(c *gin.Context) {
	var buf v1.CheckRules
	var out v1.ScriptOut

	_ = json.Unmarshal(static.CHECK_RULES, &buf)

	clone := buf

	createEnvcFile()

	for idx, rule := range buf.Rules {
		var messages = []string{}
		var checkStatus = "OK"
		o, err := exec.Command("bash", "/tmp/env_check.sh", rule.Func).Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		_ = json.Unmarshal(o, &out)
		if out.Status == "FAILED" {
			checkStatus = "FAILED"
			messages = append(messages, strings.Split(c.Request.Host, ":")[0]+`: `+out.Message)
		}
		message := utils.ToString(messages, ",")
		clone.Rules[idx].Message = message
		clone.Rules[idx].Status = checkStatus
	}

	core.WriteResponse(c, nil, &clone.Rules)
}

func createEnvcFile() {
	if utils.IsExist("/tmp/env_check.sh") {
		return
	}
	file, err := os.OpenFile("/tmp/env_check.sh", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = io.WriteString(file, string(static.ENVCHECK))
	if err != nil {
		return
	}
	return
}
