package user

import (
	"errors"
)

type User struct {
	ID, Username, Password string
}

type UserRepository interface {
	FindByUsername(string)
	Save(User)
}

type UserMemory map[string]*User

type UserRepositoryInMemory struct {
	Store UserMemory
}

func (usrRep UserRepositoryInMemory) FindByUsername(username string) (*User, error) {
	for _, u := range usrRep.Store {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (usrRep UserRepositoryInMemory) Save(user *User) {
	usrRep.Store[user.ID] = user
}
