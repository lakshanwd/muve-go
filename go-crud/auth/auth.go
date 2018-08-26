package auth

import (
	"net/http"
	"time"

	"github.com/gbrlsnchs/jwt"
	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/muve-go/go-crud/dao"
)

var signer jwt.Signer

//JWT - auth middlewear
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authHeader, err := jwt.FromRequest(c.Request); err == nil {
			if err = authHeader.Verify(signer); err == nil {
				c.Next()
			} else {
				now := time.Now()
				algValidator := jwt.AlgorithmValidator(jwt.MethodHS256)
				expValidator := jwt.ExpirationTimeValidator(now)
				if err = authHeader.Validate(algValidator, expValidator); err != nil {
					switch err {
					case jwt.ErrAlgorithmMismatch:
						c.AbortWithError(http.StatusBadRequest, err)
					case jwt.ErrTokenExpired:
						c.AbortWithError(http.StatusUnauthorized, err)
					}
				} else {
					c.Next()
				}
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

//GetAuthToken - returns JWT
func GetAuthToken(user *dao.User) (string, error) {
	// Set the options.
	now := time.Now()
	opt := &jwt.Options{
		Timestamp:      true,
		ExpirationTime: now.Add(15 * time.Second),
		NotBefore:      now.Add(30 * time.Second),
		Issuer:         "auth_server",
		Public:         map[string]interface{}{"Role": user.Role, "Name": user.Name, "UserID": user.UserID},
	}

	// Issue a new token.
	return jwt.Sign(signer, opt)
}

//InitAuth - setup auth signer
func InitAuth() {
	signer = jwt.HS256("go-practical-test")
}
