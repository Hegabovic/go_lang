package structs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func NewUser(fName string, lName string, dob string) *User {
	user := User{
		FirstName: fName,
		LastName:  lName,
		BirthDate: dob,
		CreatedAt: time.Now(),
	}
	return &user
}

// OutPutUserDetails receiver argument : go feature connect the function with the struct
func (user *User) OutPutUserDetails() {
	//(*user).FirstName
	fmt.Printf("My name is %v %v (born on %v)", user.FirstName, user.LastName, user.BirthDate)
}

func GetData(promptText string) (cleanInput string) {
	fmt.Printf(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanInput = strings.Replace(userInput, "\n", "", -1)
	return
}
