package main
import "fmt"

//structure
type Student struct {
	id int
	add string
	age int
	dpt string
}
//age increment function using pointer
func (x *Student) incAge(val int){
	x.age = x.age + val
}

func displayInfo(s Student){
	// fmt.Println(s.id)
	// fmt.Println(s.add)
	// fmt.Println(s.age)
	// fmt.Println(s.dpt)
	fmt.Printf("ID: %v, Department: %v, Age: %v, Address: %v\n",s.id,s.dpt,s.age,s.add)
}

func main() {
shykat := Student{101, "Dhaka",24,"CSE"}
ashiful := Student{102, "Dhaka",25,"CSE"}
shykat.incAge(3)
ashiful.incAge(5)
//fmt.Println(shykat)
//fmt.Println(ashiful)
// fmt.Println(shykat.id)
// fmt.Println(ashiful.id)
// fmt.Println(shykat.dpt)
// fmt.Println(ashiful.dpt)
fmt.Println("Shykat Information")
displayInfo(shykat)
fmt.Println("Ashiful Information")
displayInfo(ashiful)

}