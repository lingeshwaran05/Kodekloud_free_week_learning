package main
import "fmt"
func add(a int,b int)int{
	sum:=a+b
	return sum
}
func greet(str string){
	fmt.Print("hello ",str)
}
func main(){

	//functions in go
	
	c:=5
	d:=10
	add(c,d)
	fmt.Print("sum of ",c," and ",d," is ",add(c,d))
	greet("john")
	greet("lingesh")
	//multi return type
	// func <function name>(<parameters>)<return type>{}
	fmt.Print(addsub(5,10))
	//variadic function
	// func <function name>(<parameters>...)<return type>{}
	fmt.Print(sum(1,2,3,4,5,6,7,8,9,10))
	sub("john","maths","science","english")
	sub("lingesh","maths","science","english","social")
	
}
func sub(name string,subject ...string){
	fmt.Println(name)
	for _,v:=range subject{
		fmt.Println(v)
	}
}

func sum(a ...int)int{
	sum:=0
	for _,v:=range a{
		sum+=v
	}
	return sum
}
func addsub(a int,b int)(int,int){
	sum:=a+b
	sub:=a-b
	return sum,sub
}