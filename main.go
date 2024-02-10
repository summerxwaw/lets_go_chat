package main

import (
	"fmt"

	"github.com/summerxwaw/lets_go_chat/message"
	"github.com/summerxwaw/lets_go_chat/user"
)

func main() {

	userRepo := user.UserRepositoryInMemory{Store: user.UserMemory{}}
	msgRepo := message.MessageRepositoryInMemory{Store: message.MessageMemory{}}

	user := user.User{ID: "1", Username: "GoLEARNING", Password: "password"}
	userRepo.Save(&user)

	foundUser, err := userRepo.FindByUsername("GoLEARNING")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(foundUser)

	message := message.Message{ID: "1", Text: "Hello, GO"}
	msgRepo.Save(&message)

	if err != nil {
		fmt.Println(err)
	}

	foundMsgs, err := msgRepo.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(foundMsgs)
}
