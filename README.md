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

## License
This project is licensed under the Unlicense.