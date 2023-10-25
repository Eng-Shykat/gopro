package main

import "fmt"
type User struct { //create a struct type and store in User variable
	Name string
	Email string
	Age int
	Status bool
}
func (u User) GetStatus(){ //method GetStatus() of User struct type variable u 
	fmt.Println("Is user active?", u.Status)
}
func (u User)NewMail(){
	u.Email = "test@gmail.go"
	fmt.Println("Email of user is changed to", u.Email)
}
func (u User)UpadteAge(){
	u.Age = 29
	fmt.Println("Age of user is changed to", u.Age)
}

func main() {
	fmt.Println("Welcome to method Study!")
	//no inheritance in Go; No super or base keyword in Go; No extends keyword in Go

	shykat := User{"Shykat", "shykat@gmail.com", 26, true}
	fmt.Println(shykat)
	fmt.Printf("Details of shykat is %+v\n", shykat)
	shykat.GetStatus() //calling method GetStatus() of User struct type variable shykat 
	shykat.NewMail() //calling method NewMail() of User struct type variable shykat
	shykat.UpadteAge() //calling method UpadteAge() of User struct type variable shykat
	fmt.Printf("Update Details of shykat is %+v\n", shykat)
}