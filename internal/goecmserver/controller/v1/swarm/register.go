package swarm

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"net/http"
	"os"
)

//nolint: not used.
func (s *SwarmController) Register(c *gin.Context) {
	log.L(c).Debug("host register function called.")
	var host v1.HostInfo
	host.HostIP = c.ClientIP()
	err := s.svc.Swarm().Register(c, &host, metav1.CreateOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "数据库错误"), nil)
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=add_host_key.sh")
	c.Header("Content-Transfer-Encoding", "application/text/plain")
	c.Header("Cache-Control", "no-cache")

	content, _ := os.ReadFile(utils.UserHome() + "/.ssh/add_host_key.sh")

	c.Status(http.StatusOK)
	c.Writer.Write(content)
}
