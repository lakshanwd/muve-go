package repo

import (
	"testing"

	"github.com/lakshanwd/muve-go/go-crud/dao"
)

func TestAuthenticate(t *testing.T) {
	cases := []struct {
		userName       string
		password       string
		expectingValue bool
	}{
		{"admin", "123", true},
		{"developer", "123", true},
		{"user", "123", true},
		{"root", "123", false},
		{"admin", "1234", false},
	}
	for _, testCase := range cases {
		result := Authenticate(&dao.UserCredentials{Username: testCase.userName, Password: testCase.password})
		if (result != nil) != testCase.expectingValue {
			t.Errorf("username=%v, password=%v and got a result=%v", testCase.userName, testCase.password, testCase.expectingValue)
		}
	}
}

func TestGetBookings(t *testing.T) {
	cases := []struct {
		from  int64
		count int
	}{
		{0, 1000},
		{0, 10000},
	}
	for _, testCase := range cases {
		if list, err := GetBookings(testCase.from, testCase.count); err != nil {
			t.Errorf("%v", err.Error())
		} else if list.Len() != testCase.count {
			t.Errorf("got %v rows, expected row count is %v", list.Len(), testCase.count)
		}
	}
}
