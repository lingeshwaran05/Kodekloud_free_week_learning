package main
import "fmt"
func main() {
//pointer var <pointer>*<datatype>
x:=10
//<pointer name>:=&<variable name>
var p *int=&x
pointer:=&x

fmt.Print(pointer,*pointer)
fmt.Println(p)
//adress of x
fmt.Println(*p)
//deferencing
*p=20
fmt.Println(x)
fmt.Println("Adress %Tand defernce%v",&x,*(&x))

}