package user

import (
	"errors"
	"sync"
)

type User struct {
	ID, Username, Password string
}

type Repository interface {
	FindByUsername(string)
	Save(User)
}

type Memory map[string]User

type UserRepository struct {
	store Memory
	mutex *sync.RWMutex
}

func NewUserRepository() UserRepository {
	return UserRepository{
		store: make(Memory),
		mutex: &sync.RWMutex{},
	}
}

func (u UserRepository) FindByUsername(username string) (User, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	for _, u := range u.store {
		if u.Username == username {
			return u, nil
		}
	}

	return User{}, errors.New("user not found")
}

func (u UserRepository) Save(user User) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	u.store[user.ID] = user
}
