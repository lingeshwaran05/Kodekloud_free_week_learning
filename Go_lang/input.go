package main
import "fmt"
"reflect"
func main(){
	var name string	
	var is bool
	fmt.Print("Enter name & is: ")
	fmt.Scanf("%s %t",&name,&is)
	fmt.Println(name,is)
	fmt.Printf("type of name is %T",name)
	fmt.Printf("type of is is %T",is)
	fmt.Printf("Type:%v",reflect.TypeOf(name))
	fmt.Printf("Type:%v",reflect.TypeOf(is))
}