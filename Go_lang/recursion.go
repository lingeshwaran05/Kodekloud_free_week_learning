package main
import "fmt"
func fact(n int)int{
	if n==0{
		return 1
	}
	return n*fact(n-1)
	//higher order function
	
}
func result(radius float64,cal func(r float64)float64){
	fmt.Println(cal(radius))
	fmt.Println("THank you for using our service")
}
func cal(r float64)float64{
	return 3.14*r*r
}
func main(){
	fmt.Print(fact(5))
	result(5,cal)
	result(10,cal)
}