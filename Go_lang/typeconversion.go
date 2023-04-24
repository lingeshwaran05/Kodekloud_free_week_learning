package main
import "fmt"
"strconv"
func main(){
	//int to float
	var a int=90
	var b float64=float64(a)
	fmt.Println(b)
	//float to int
	var c float64=90.90
	var d int=int(c)
	fmt.Println(d)
	//string conversion int to string
	var e int=90
	var f string=strconv.Itoa(e)
	fmt.Println(f)
	//string conversion string to int
	var g string="90"
	var h int,_=strconv.Atoi(g)
}