package repo

import (
	"strings"

	"github.com/lakshanwd/muve-go/go-crud/dao"
)

//Authenticate - authenticate user
func Authenticate(credentials *dao.UserCredentials) (authenticatedUser *dao.User) {
	userList := []struct {
		userName string
		password string
		user     *dao.User
	}{
		{"admin", "123", &dao.User{UserID: 1, Name: "Administrator", Role: "admin"}},
		{"developer", "123", &dao.User{UserID: 2, Name: "Lakshan Dissanayake", Role: "developer"}},
		{"user", "123", &dao.User{UserID: 3, Name: "Rob Brew", Role: "user"}},
	}
	for _, value := range userList {
		if strings.Compare(value.userName, credentials.Username) == 0 && strings.Compare(value.password, credentials.Password) == 0 {
			authenticatedUser = value.user
			return
		}
	}
	return
}
