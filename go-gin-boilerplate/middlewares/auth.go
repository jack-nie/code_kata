package middlewares

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jack-nie/code_kata/go-gin-boilerplate/config"
	"github.com/sirupsen/logrus"
)

func Auth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		token, err := jwt.Parse(authorization, secret())
		if err != nil {
			logrus.Info("error while parsing token: %v", err)
			c.AbortWithStatus(400)
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logrus.Info("error while convert claim to mapclaim")
			c.AbortWithStatus(400)
		}

		if !token.Valid {
			logrus.Info("token is unvalid")
			c.AbortWithStatus(401)
		} else {
			if role == "admin" {
				if adminID, ok := claim["AdminID"]; !ok {
					logrus.Info("not valid admin")
					c.AbortWithStatus(401)
				} else {
					c.Set("AdminID", adminID)
				}
			}

			if role == "user" {
				if userID, ok := claim["UserID"]; !ok {
					logrus.Info("not valid user")
					c.AbortWithStatus(401)
				} else {
					c.Set("UserID", userID)
				}
			}
		}
		c.Next()
	}
}

func secret() jwt.Keyfunc {
	c := config.GetConfig()
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(c.GetString("jwt.secret")), nil
	}
}
