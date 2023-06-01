package swarm

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	myssh "go-ecm/utils/ssh"
	"strconv"
)

func (s *SwarmController) Check(c *gin.Context) {
	log.L(c).Debug("host check function called.")
	var connectionInfo v1.HostConnectionCheck

	if err := c.ShouldBindJSON(&connectionInfo); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "客户端传参错误"), nil)
		return
	}

	err := check(&connectionInfo)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrSSHConnectionFailed, "无法与目标主机建立链接"), nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}

func check(data *v1.HostConnectionCheck) error {
	port, _ := strconv.Atoi(data.SSHPort)
	_, err := myssh.NewClient(data.SSHUsername,
		myssh.WithAuthByKey(utils.UserHome()+"/.ssh/id_rsa"),
		myssh.WithPort(port),
		myssh.WithHost(data.HostIPAddress))

	if err != nil {
		return err
	}

	return nil
}
