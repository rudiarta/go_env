package service

type User struct {
	Name string
	Age int
}

func (u *User) GetUser() string {
	return u.Name
}

func Ambil(u User) User {
	return u
}