package main
import "fmt"
func main(){
	fmt.Println("like C++")
	fmt.Println("Go is a statically typed language")
	fmt.Println("Go is a some times feels like Dynamically typed language")
	name:="This is string"
	fmt.Println(name)
	integer:=90
	fmt.Println(integer)
	float:=90.90
	fmt.Println(float)
	complex:=90+90i
	fmt.Println(complex)
	array:=[]int{1,2,3,4,5}
	fmt.Println(array)
	boolean:=true
	fmt.Println(boolean)
	mapping:=map[string]int{"one":1,"two":2}
	fmt.Println(mapping)
	fmt.Printf("THis if for string %v",name)
	fmt.Printf("THis if for decimal %d",integer)
	fmt.Printf("THis if for type %T",float)
	//declaring with or without var keyboard
	var (s string="foo"
	i int=5)
	fmt.Print(s,i)
	//this is short declaration
	str:="bar"
	fmt.Print(str)
}