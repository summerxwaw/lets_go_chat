package message

import (
	"errors"
)

type Message struct {
	ID, Text string
}
type MessageRepository interface {
	FindAll()
	Save(Message)
}

type MessageMemory map[string]*Message

type MessageRepositoryInMemory struct {
	Store MessageMemory
}

func (msgRep *MessageRepositoryInMemory) FindAll() (map[string]*Message, error) {
	return msgRep.Store, errors.New("messages not found :(")
}

func (msgRep *MessageRepositoryInMemory) Save(msg *Message) error {
	msgRep.Store[msg.ID] = msg

	return nil
}
