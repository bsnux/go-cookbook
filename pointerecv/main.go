package main

/**
* Pointers receivers explained with simple examples
 */

import (
	"fmt"
)

type School struct {
	name string
}

type User struct {
	age  int
	name string
}

// Student is using composition
type Student struct {
	school School
	user   User
}

// changeAge is using a pointer receiver therefore original struct will be
// modified
func (user *User) changeAge(age int) {
	user.age = age
}

// displayAge is not modifying any value then we're not using a pointer receiver
func (user User) displayAge() {
	fmt.Println(user.age)
}

func main() {
	fmt.Println("Hello, playground")

	user := User{23, "Lucas"}

	fmt.Println(user)

	user.age = 25
	fmt.Println(user)
	fmt.Println(user.age)

	user.changeAge(30)
	fmt.Println(user)

	user.displayAge()

	school := School{"USC"}

	student := Student{school, user}
	fmt.Println(student)

	fmt.Println(student.user.name)

}
