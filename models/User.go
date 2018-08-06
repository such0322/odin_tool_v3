package models

import (
	"errors"
	"time"

	"github.com/go-macaron/session"
)

type User struct {
	ID        int
	Account   string
	Password  string
	Name      string
	Status    int
	CreatedAt int64
}

func NewUser() *User {
	return &User{}
}

func GetUserById(id int) *User {
	var user User
	db.First(&user, id)
	return &user
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
