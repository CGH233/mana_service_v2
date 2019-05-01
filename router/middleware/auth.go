package middleware

import (
	"fmt"

	b64 "encoding/base64"
	s "strings"

	"github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		b, err := b64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			handler.SendUnauthorized(c, errno.ErrTokenInvalid, nil, err.Error())
			c.Abort()
			return
		}

		arr := s.Split(string(b), ":")

		if len(arr) == 2 && arr[0] == viper.GetString("admin.name") && arr[1] == viper.GetString("admin.pwd") {
			fmt.Println(arr[0], arr[1])
			c.Next()
		} else {
			handler.SendUnauthorized(c, errno.ErrTokenInvalid, nil, err.Error())
			c.Abort()
			return
		}
	}
}
