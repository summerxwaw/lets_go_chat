## How to use

This package exports two functions for hashing a password and to check whether a password matches a hash:

  ```go
   HashPassword(password string) (string, error)
   ```
  ```go
   CheckPasswordHash(password, hash string) bool
   ```

You can import the hasher package and call these two functions with proper arguments as in this example:
 ```go
 package main
  
import (
  "fmt"
  "lets-go-chat/pkg/hasher"
)
  
func main() {
  password := "your_password"

  hash, err := hasher.HashPassword(password)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  fmt.Println(hasher.CheckPasswordHash(password, hash))
}
   ```

For using Meessage and User repository we must follow this example:
  ```go
  import (
	"github.com/summerxwaw/lets_go_chat/message"
	"github.com/summerxwaw/lets_go_chat/user"
)

	userRepo := user.NewUserRepository()
	msgRepo := message.NewMessageRepository()

	user := user.User{ID: "1", Username: "GoLEARNING", Password: "password"}
	userRepo.Save(user)

	foundUser, err := userRepo.FindByUsername("GoLEARNING")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(foundUser)

	message := message.Message{ID: "1", Text: "Hello, GO"}
	msgRepo.Save(message)

	if err != nil {
		fmt.Println(err)
	}

	foundMsgs, err := msgRepo.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(foundMsgs)

   ```
## License
This project is licensed under the Unlicense.