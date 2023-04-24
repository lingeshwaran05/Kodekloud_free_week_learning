package main
import "fmt"
func main(){
	a:=10
	b:=20
	//simple if
	if a>b{
		fmt.Println("a is greater than b")
	}
	//if else
	if a<b{
		fmt.Println("a is less than b")
	}else{
		fmt.Println("a is greater than b")
	}
//if else if

if a<b{
	fmt.Println("a is less than b")
}else if b>a{
	fmt.Println("b is greater than a")
}else{
	fmt.Println("a is equal to b")
}


//nested if
	var c string="fruit"
	var d string="apple"
	if c=="fruit"{
		if d=="apple"{
			fmt.Println("apple is a fruit")
		}
	}


//switch case
switch(1){
case 1:
	fmt.Println("one")
	break
case 2:
	fmt.Println("two")
	break
case 3:
	fmt.Println("three")
	break
	default:
	fmt.Println("default")
	break
}
	//switch case with multiple cases
	switch(1){
	case 1,2,3:
		fmt.Println("one")
		//break can be implicit
	case 4,5,6:
		fmt.Println("two")
		default:
		fmt.Println("default")
	}
	//swtich case with no condition
	var e int=10
	switch{
	case e<10:
		fmt.Println("e is less than 10")
	case e>10:
		fmt.Println("e is greater than 10")
	default:
		fmt.Println("e is equal to 10")
	}
}