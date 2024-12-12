package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status string
}

func main() {
	fmt.Println("Methds in go")
	user := User{Name: "John Doe", Age: 30, Email: "john@example.com", Status: "Active"}
	user.GetStatus()
	user.NewEmail()
	fmt.Println(user) // we will get the old emailonly it craeted the copy and do chnages in it
	// if wnats to chage the origin al value then we needs to craete the pointers
}

func (u User) GetStatus() {
	fmt.Printf("User %s is %s\n", u.Name, u.Status)

}

// we can maipulate the propeties in the stuct
func (u User) NewEmail() {
	u.Email = "new@" + u.Email
	fmt.Printf("New email for user %s is %s\n", u.Name, u.Email)
}
