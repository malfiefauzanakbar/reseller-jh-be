package common

import (
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
		if authorization == "" {
			base.RespondError(c, http.StatusUnauthorized, constant.Unauthorized, "Missing Authorization header")
			c.Abort()
			return
		}

		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			base.RespondError(c, http.StatusUnauthorized, constant.Unauthorized, "Invalid Authorization format")
			c.Abort()
			return
		}

		token := parts[1]
		decryptToken, err := Decrypt(token)
		if err != nil {
			base.RespondError(c, http.StatusUnauthorized, constant.Unauthorized, err)
			c.Abort()
			return
		}

		session := sessions.Default(c)
		if session.Get("token") != decryptToken {
			base.RespondError(c, http.StatusUnauthorized, constant.Unauthorized, "Token mismatch")
			c.Abort()
			return
		}

		now := time.Now()
		dateTime := now.Format("2006-01-02 15:04:05")

		lastAccessed, ok := session.Get("lastAccessed").(string)
		if !ok || lastAccessed == "" {
			session.Set("lastAccessed", dateTime)
			if err := session.Save(); err != nil {
				base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err)
				c.Abort()
				return
			}
			c.Next()
			return
		}

		localLocation := time.Now().Location()
		lastAccessedTime, err := time.ParseInLocation("2006-01-02 15:04:05", lastAccessed, localLocation)
		if err != nil {
			base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err)
			c.Abort()
			return
		}

		if now.Sub(lastAccessedTime) >= time.Hour {
			session.Clear()
			base.RespondError(c, 440, constant.SessionExpired, "Session expired")
			c.Abort()
			return
		}

		session.Set("lastAccessed", dateTime)
		if err := session.Save(); err != nil {
			base.RespondError(c, http.StatusInternalServerError, constant.InternalServerError, err)
			c.Abort()
			return
		}

		c.Next()
	}
}
