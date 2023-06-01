package menu

import (
	"github.com/gin-gonic/gin"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (m *MenuController) Get(c *gin.Context) {
	log.Debug("Menu get fucntion called.")
	var menu []v1.MenuBaseSpec
	menu, err := m.srv.Menu().Get(c, metav1.ListOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
	}

	var menus = &v1.Menus{
		Menus: menu,
	}
	core.WriteResponse(c, nil, menus)
}
