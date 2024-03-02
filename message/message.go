package message

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Message struct {
	ID, Text string
}
type Repository interface {
	FindAll()
	Save(Message)
}

func (m *MessageRepository) Find(id string) (Message, error) {
	filter := bson.M{"id": id}
	var message Message

	err := m.collection.FindOne(context.Background(), filter).Decode(&message)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}

func (m *MessageRepository) Save(message Message) error {
	_, err := m.collection.InsertOne(context.Background(), message)
	return err
}
