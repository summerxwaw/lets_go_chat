package user_test

import (
	"testing"

	"github.com/summerxwaw/lets_go_chat/user"
)

func TestUserRepository(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		expectedErr string
		user        user.User
	}{
		{
			name:     "user exists",
			username: "existinguser",
			user:     user.User{ID: "1", Username: "existinguser", Password: "pass"},
		},
		{
			name:        "user not found",
			username:    "nonexistentuser",
			expectedErr: "user not found",
		},
	}

	repo := user.NewUserRepository()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.user.ID != "" {
				repo.Save(test.user)
			}

			_, err := repo.FindByUsername(test.username)

			if err != nil && err.Error() != test.expectedErr {
				t.Errorf("Expected error %q but got %q", test.expectedErr, err)
			} else if err == nil && test.expectedErr != "" {
				t.Errorf("Expected error %q but got nil", test.expectedErr)
			}
		})
	}
}
