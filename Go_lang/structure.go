package main
import "fmt"
type student struct{
	name string
	rollno int
	marks []int

}
func main(){
//structures := struct {}
//type structname struct{}
//var name struct
var s student 
fmt.Print(s)
s.name = "Raj"
s.rollno = 1
s.marks = []int{1,2,3,4,5}
fmt.Println(s)
//variables := new(struct name)
st:=new(student)
st.name = "Raj"
st.rollno = 1
st.marks = []int{1,2,3,4,5}
fmt.Println(st)
stu:=student{"lingesh",2,[]int{1,2,3,4,5}}

fmt.Println(stu)


}