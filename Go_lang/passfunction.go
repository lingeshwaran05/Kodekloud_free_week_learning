package main
import "fmt"
//pass by value
func modify(a int) {
	a = 10
	fmt.Println("Inside modify function ", a)
}
//pass by reference
func modify1(s *string) {
	*s="modified"
	fmt.Println("Inside modify function ", *s)
}
func main() {
//pass by value
a:=50
fmt.Println("Before function call ", a)
modify(a)
fmt.Println("After function Inside main function ", a)
h:="original"
fmt.Print(h)
modify1(&h)
fmt.Print(h)

}