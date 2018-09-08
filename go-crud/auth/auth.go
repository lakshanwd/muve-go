package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt"
	"github.com/gin-gonic/gin"
	"github.com/lakshanwd/muve-go/go-crud/dao"
)

var signer jwt.Signer

//Token - Auth token structure
type Token struct {
	*jwt.JWT
	Role   string `json:"role,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int64  `json:"userID,omitempty"`
}

//JWT - auth middlewear
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, "Bearer ")[1]
		if payload, sig, err := jwt.Parse(token); err == nil {
			var opt Token
			if err = jwt.Unmarshal(payload, &opt); err == nil {
				if err = signer.Verify(payload, sig); err == nil {
					now := time.Now()
					issuedAt := jwt.IssuedAtValidator(now)
					expired := jwt.ExpirationTimeValidator(now)
					audience := jwt.AudienceValidator("muve")
					notBefore := jwt.NotBeforeValidator(now)
					issuer := jwt.IssuerValidator("auth_server")
					iD := jwt.IDValidator("unique-id")
					subject := jwt.SubjectValidator("developer")
					if err = opt.Validate(issuedAt, expired, audience, notBefore, issuer, iD, subject); err != nil {
						switch err {
						case jwt.ErrIatValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "issuedAt validation failed"})
						case jwt.ErrExpValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "expired validation failed"})
						case jwt.ErrAudValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "audience validation failed"})
						case jwt.ErrNbfValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "notBefore validation failed"})
						case jwt.ErrIssValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "issuer validation failed"})
						case jwt.ErrJtiValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "iD validation failed"})
						case jwt.ErrSubValidation:
							c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "subject validation failed"})
						}
					} else {
						c.Next()
					}
				} else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "verification failed"})
				}
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "invalid token"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "unable to parse token"})
		}
	}
}

//GetAuthToken - returns JWT
func GetAuthToken(user *dao.User) (token string, err error) {
	// Set the options.
	now := time.Now()
	opt := &Token{
		JWT: &jwt.JWT{
			Issuer:         "auth_server",
			Subject:        "developer",
			Audience:       "muve",
			ExpirationTime: now.Add(30 * time.Minute).Unix(),
			NotBefore:      now.Add(10 * time.Minute).Unix(),
			IssuedAt:       now.Unix(),
			ID:             "unique-id",
		},
		Role:   user.Role,
		Name:   user.Name,
		UserID: user.UserID,
	}
	opt.SetAlgorithm(signer)
	opt.SetKeyID("auth-token")

	if payload, err := jwt.Marshal(opt); err == nil {
		if bytes, err := signer.Sign(payload); err == nil {
			token = fmt.Sprintf("%s", bytes)
		}
	}
	return
}

//InitAuth - setup auth signer
func InitAuth() {
	signer = jwt.NewHS256("go-practical-test")
}
