package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/store"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"net/http"
)

func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := isAdmin(c); err != nil {
			switch c.FullPath() {
			case "/v1/users":
				if c.Request.Method != http.MethodPost {
					core.WriteResponse(c, errors.WithCode(code.ErrPermissionDenied, ""), nil)
					c.Abort()

					return
				}
			case "/v1/users/:name", "/v1/users/:name/change_password":
				username := c.GetString("username")
				if c.Request.Method == http.MethodDelete ||
					(c.Request.Method != http.MethodDelete && username != c.Param("name")) {
					core.WriteResponse(c, errors.WithCode(code.ErrPermissionDenied, ""), nil)

					c.Abort()
					return
				}

			default:
			}

			c.Next()
		}
	}
}

func isAdmin(c *gin.Context) error {
	username := c.GetString("username")

	user, err := store.Client().Users().Get(c, username, metav1.GetOptions{})
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	if user.Name != "admin" {
		return errors.WithCode(code.ErrPermissionDenied, "user %s is not a administrator", user.UserName)
	}

	return nil
}
