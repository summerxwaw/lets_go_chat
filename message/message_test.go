package message

import (
	"testing"
)

func TestMessageRepositoryInMemory_FindAll(t *testing.T) {
	repo := MessageRepositoryInMemory{
		Store: MessageMemory{
			"id1": &Message{ID: "id1", Text: "Hello, Go!"},
		},
	}

	msgs, err := repo.FindAll()
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if msgs == nil || len(msgs) != 1 || msgs["id1"].Text != "Hello, Go!" {
		t.Errorf("Got unexpected messages: %v", msgs)
	}

	repoEmpty := MessageRepositoryInMemory{Store: MessageMemory{}}
	_, err = repoEmpty.FindAll()
	if err == nil {
		t.Errorf("Expected error for empty repository")
	}
}

func TestMessageRepositoryInMemory_Save(t *testing.T) {
	repo := MessageRepositoryInMemory{Store: MessageMemory{}}

	msg := Message{ID: "id1", Text: "Hello, Go!"}
	repo.Save(&msg)

	if len(repo.Store) != 1 {
		t.Errorf("Expected one message in the store")
	}
	if repo.Store[msg.ID].Text != "Hello, Go!" {
		t.Errorf("Got unexpected message: %v", repo.Store[msg.ID])
	}
}
