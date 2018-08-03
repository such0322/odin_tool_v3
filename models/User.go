package models

import (
	"errors"
	"time"

	"github.com/go-macaron/session"
)

type User struct {
	ID        int64
	Account   string
	Password  string
	Name      string
	Status    int64
	CreatedAt int64
}

func GetTestUser() *User {
	t := time.Now()
	return &User{
		ID:        1,
		Account:   "moz1",
		Password:  "",
		Name:      "moz_name",
		Status:    1,
		CreatedAt: t.Unix(),
	}
}

func UserSignin(sess session.Store) *User {
	uid := sess.Get("uid")
	if uid == nil {
		return nil
	}
	user := GetTestUser()
	return user
}

func UserLogin(account, passwd string) (*User, error) {
	t := time.Now()
	if account == "moz1" {
		return &User{
			ID:        1,
			Account:   "moz1",
			Password:  "",
			Name:      "moz_name",
			Status:    1,
			CreatedAt: t.Unix(),
		}, nil
	}
	return nil, errors.New("不是moz1")
}

func getUserByAccount(account string) (*User, error) {
	return nil, nil
}
