package main
import "fmt"
func main(){
	var a=10
	var b=20
	//equality
	fmt.Println(a==b)
	fmt.Println(a != b)
	//relational
	fmt.Println(a<b)
	fmt.Println(a>b)
	fmt.Println(a<=b)
	fmt.Println(a>=b)
	//arithmetic
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	c:=1
	c++
	a--
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(a%b)
	//assignment
	a+=b
	fmt.Println(a)
	a=100
	a-=b
	fmt.Println(a)
	//logical
	a=10
	b=10

	fmt.Println((a<b)&&(a==b))
	fmt.Println((a==b)||(b<=a))
	fmt.Println(!true)
	fmt.Println(a&b)
	fmt.Println(a|b)
	fmt.Println(a^b)
	//bitwise
	fmt.Println(a&^b)
	//shift
	fmt.Println(a<<b)
	fmt.Println(a>>b)

}