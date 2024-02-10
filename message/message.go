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

func (msgRep MessageRepositoryInMemory) FindAll() (MessageMemory, error) {
	if msgRep.Store == nil {
		return nil, errors.New("messages not found :(")
	}
	return msgRep.Store, nil
}

func (msgRep MessageRepositoryInMemory) Save(msg *Message) {
	msgRep.Store[msg.ID] = msg
}
