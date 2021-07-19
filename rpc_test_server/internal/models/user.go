package models

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
