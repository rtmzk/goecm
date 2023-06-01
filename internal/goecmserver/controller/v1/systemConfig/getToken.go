package systemConfig

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (scc *SystemConfigController) GetToken(c *gin.Context) {
	log.L(c).Debug("get agent token function called.")

	data, err := scc.srv.SystemConfig().GetAgentToken(c, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, data)
}
