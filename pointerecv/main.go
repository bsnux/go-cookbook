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
	user := User{23, "Lucas"}
	fmt.Println(user)
	user.age = 25
	fmt.Println(user)
	fmt.Println(user.age)

	user.changeAge(30)
	fmt.Println(user)
	// Following line is equivalent to `user.changeAge(50)`, Go does the work for us
	(&user).changeAge(50)
	fmt.Println(user)
	user.displayAge()

	bill := &User{40, "Bill"}
	fmt.Println(bill.name)

	bill.changeAge(44)
	fmt.Println(bill.age)
	// Following line is equivalent to `bill.changeAge(55)`, Go does the work for us
	(*bill).changeAge(55)
	fmt.Println(bill.age)

	// Composition example,
	school := School{"USC"}
	student := Student{school, user}
	fmt.Println(student)
	fmt.Println(student.user.name)
}
