package repo

import (
	"strings"

	"github.com/lakshanwd/muve-go/go-crud/dao"
)

//Authenticate - authenticate user
func Authenticate(credentials *dao.UserCredentials) (authenticatedUser *dao.User) {
	type userCredential struct {
		UserName string
		Password string
		User     *dao.User
	}
	userList := []*userCredential{
		&userCredential{UserName: "admin", Password: "123", User: &dao.User{UserID: 1, Name: "Administrator", Role: "admin"}},
		&userCredential{UserName: "developer", Password: "123", User: &dao.User{UserID: 2, Name: "Lakshan Dissanayake", Role: "developer"}},
		&userCredential{UserName: "user", Password: "123", User: &dao.User{UserID: 3, Name: "Rob Brew", Role: "user"}},
	}
	for _, value := range userList {
		if strings.Compare(value.UserName, credentials.Username) == 0 && strings.Compare(value.Password, credentials.Password) == 0 {
			authenticatedUser = value.User
			return
		}
	}
	return
}
