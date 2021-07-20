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

func (u *User) GetFilter(where ...interface{}) string {
	var filter string

	whereMap, ok := where[0].(map[string]interface{})
	if !ok {
		return filter
	}

	if len(whereMap) == 0 {
		return filter
	}

	filter += ` WHERE `
	for key, value := range whereMap {
		vString, ok := value.(string)
		if ok {
			filter += "AND " + key + " = " + fmt.Sprintf("'%s'", vString)
		} else {
			filter += "AND " + key + " = " + fmt.Sprintf("%v", value)
		}
	}

	return filter
}
