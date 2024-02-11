package message

import (
	"errors"
	"sync"
)

type Message struct {
	ID, Text string
}
type Repository interface {
	FindAll()
	Save(Message)
}

type Memory map[string]Message

type MessageRepository struct {
	store Memory
	mutex *sync.RWMutex
}

func NewMessageRepository() MessageRepository {
	return MessageRepository{
		store: make(Memory),
		mutex: &sync.RWMutex{},
	}
}

func (m MessageRepository) FindAll() (Memory, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.store == nil || len(m.store) == 0 {
		return Memory{}, errors.New("messages not found :(")
	}

	return m.store, nil
}

func (m MessageRepository) Save(msg Message) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.store[msg.ID] = msg
}
