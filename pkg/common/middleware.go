package common

import (
	"fmt"
	"net/http"
	"reseller-jh-be/base"
	"reseller-jh-be/constant"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		parts := strings.Split(authorization, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			base.RespondError(c, http.StatusUnauthorized, constant.Unauthorized, nil)
			c.Abort()
			return
		}

		token := parts[1]
		decryptToken, err := Decrypt(token)
		if err != nil {
			base.RespondError(c, http.StatusInternalServerError, constant.Unauthorized, err)
			c.Abort()
			return
		}

		session := sessions.Default(c)
		if token == "" || decryptToken != session.Get("token") {
			base.RespondError(c, http.StatusUnauthorized, constant.Unauthorized, nil)
			c.Abort()
			return
		}

		now := time.Now()
		dateTime := now.Format("2006-01-02 15:04:05")
		if session.Get("lastAccessed") == nil {
			session.Set("lastAccessed", dateTime)
			session.Save()
		} else {
			lastAccessed, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v", session.Get("lastAccessed")))
			if err != nil {
				fmt.Println("Error parsing dateTime:", err)
				return
			}
			if time.Since(lastAccessed) > time.Minute {
				session.Clear()
				base.RespondError(c, 440, constant.SessionExpired, nil)
				c.Abort()
				return
			} else {
				session.Set("lastAccessed", dateTime)
				session.Save()
			}
		}

		c.Next()
	}
}
