package main
import "fmt"
func main(){
	//this is global variable outside the block
	global:="This is global variable"
	{
		//this is inside the block
		local:="This is local variable"
		fmt.Println(local)
		fmt.Println(global)
	}
	fmt.Println(global)
	//fmt.Println(local) //this will give error
}