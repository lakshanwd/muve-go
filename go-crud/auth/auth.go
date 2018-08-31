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
				now := time.Now()
				algValidator := jwt.AlgorithmValidator(jwt.MethodHS256)
				expValidator := jwt.ExpirationTimeValidator(now)
				if err = authHeader.Validate(algValidator, expValidator); err != nil {
					switch err {
					case jwt.ErrAlgorithmMismatch:
						c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"reason": "invalid algorythrm"})
					case jwt.ErrTokenExpired:
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "token expired"})
					}
				} else {
					c.Next()
				}
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "invalid token"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "No authorize header"})
		}
	}
}

//GetAuthToken - returns JWT
func GetAuthToken(user *dao.User) (string, error) {
	// Set the options.
	now := time.Now()
	opt := &jwt.Options{
		Timestamp:      true,
		ExpirationTime: now.Add(10 * time.Minute),
		NotBefore:      now.Add(15 * time.Minute),
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
