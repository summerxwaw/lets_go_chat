package message_test

import (
	"testing"

	"github.com/summerxwaw/lets_go_chat/message"
)

func TestMessageRepository(t *testing.T) {
	tests := []struct {
		name        string
		text        string
		expectedErr string
		message     message.Message
	}{
		{
			name:        "messages not found",
			expectedErr: "messages not found :(",
		},
		{
			name: "messages exists",
			text: "messages",
			message: message.Message{
				ID: "1", Text: "existingtext",
			},
		},
	}

	repo := message.NewMessageRepository()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.message.ID != "" {
				repo.Save(test.message)
			}

			_, err := repo.FindAll()

			if err != nil && err.Error() != test.expectedErr {
				t.Errorf("Expected error %q but got %q", test.expectedErr, err)
			} else if err == nil && test.expectedErr != "" {
				t.Errorf("Expected error %q but got nil", test.expectedErr)
			}

		})
	}
}
