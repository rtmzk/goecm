package user

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

// Get administrator api. Query all users info
func (u *UserController) Get(c *gin.Context) {
	log.L(c).Debug("Get user info method called.")
	user, err := u.srv.Users().Get(c, c.Param("name"), metav1.GetOptions{})

	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}

// GetCurrentUser scoped api. query current user info.
// This api used for get current user info immediately after user login successfully.
// token necessary.
func (u *UserController) GetCurrentUser(c *gin.Context) {
	log.L(c).Debug("Get current user info method called")
	user, err := u.srv.Users().Get(c, c.GetString("username"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
