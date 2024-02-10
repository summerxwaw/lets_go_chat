package user

import (
	"testing"
)

func TestUserRepositoryInMemory_FindByUsername(t *testing.T) {
	repo := UserRepositoryInMemory{
		Store: UserMemory{
			"id1": &User{ID: "id1", Username: "Username", Password: "password"},
		},
	}

	user, err := repo.FindByUsername("Username")
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if user == nil || user.ID != "id1" {
		t.Errorf("Got unexpected user: %v", user)
	}

	_, err = repo.FindByUsername("NonExistentUser")
	if err == nil {
		t.Errorf("Expected error for non existent user")
	}
}

func TestUserRepositoryInMemory_Save(t *testing.T) {
	repo := UserRepositoryInMemory{Store: UserMemory{}}

	user := User{ID: "id1", Username: "Username", Password: "password"}
	repo.Save(&user)

	if len(repo.Store) != 1 {
		t.Errorf("Expected one user in the store")
	}
	if repo.Store[user.ID].Username != "Username" {
		t.Errorf("Got unexpected user: %v", repo.Store[user.ID])
	}
}
