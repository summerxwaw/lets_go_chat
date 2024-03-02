package user

import (
	"errors"
)

type User struct {
	ID, Username, Password string
}

type Repository interface {
	FindByUsername(string)
	Save(User)
}

func (u UserRepository) FindByUsername(username string) (*User, error) {
	user := &User{}

	result := u.db.Where("username = ?", username).First(user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u UserRepository) Save(user User) error {
	if u.db == nil {
		return errors.New("database not initialized")
	}

	if user.Username == "" {
		return errors.New("User name cannot be empty")
	}

	result := u.db.Create(&UserDB{
		ID:       user.ID,
		Username: user.Username,
	})

	if result.Error != nil {
		return errors.New("error while saving user to db")
	}

	return nil
}
