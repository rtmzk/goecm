package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/constand"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
)

func (a *AgentController) Report(c *gin.Context) {
	var items v1.ReportItems
	var findFlag bool
	var findIndex int
	if err := c.ShouldBindJSON(&items); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "传参错误"), nil)
		return
	}

	if len(constand.ReportItems) == 0 {
		constand.ReportItems = append(constand.ReportItems, &items)
		core.WriteResponse(c, nil, map[string]string{"status": "recived"})
		return
	}

	for i := 0; i < len(constand.ReportItems); i++ {
		if constand.ReportItems[i].Host != items.Host {
			continue
		}
		findFlag = true
		findIndex = i
	}

	if findFlag {
		constand.ReportItems[findIndex] = &items
	} else {
		constand.ReportItems = append(constand.ReportItems, &items)
	}

	core.WriteResponse(c, nil, map[string]string{"status": "recived"})
}
