package models

import "fmt"

type User struct {
	Id        int64
	Uuid      string
	NickName  string
	Login     string
	Password  string
	Rule      int64
	Blocked   bool
	CreateAt  int64
	BlockedAt int64
	UpdateAt  int64
}

func (u User) Table() string {
	return "user"
}

func (u *User) GetFilter() string {
	var where string

	if u == nil {
		return where
	}

	if u.Id != 0 {
		where += fmt.Sprintf("id = %d", u.Id)
		return where
	}

	if len(u.Uuid) != 0 {
		where += "uuid = '" + u.Uuid + "'"
		return where
	}

	if len(u.NickName) != 0 {
		where += "nick_name = '" + u.NickName + "'"
		return where
	}

	if len(u.Login) != 0 {
		where += "login = '" + u.Login + "'"
		return where
	}

	return where
}
