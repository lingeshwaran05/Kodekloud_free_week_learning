package main
import "fmt"
func main(){
	var i int
	//for loop
for i=1; i<= 3;i++{
	fmt.Println(i*i)
}
i=1
for i<=5{
	fmt.Println(i)
	i++
}
//break loop
var j int=1
for j<=5{
	fmt.Println(j)
	if j==3{
		break
	}
	j++
	fmt.Print(j)
}
//continue loop
var k int=1
for k<=5{
	if k==3{
		k++
		continue
	}
	fmt.Println(k)
	k++
}
}