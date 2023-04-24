//var <name> map[key datatype]<data datatype>
package main
import "fmt"
func main(){
maps:=map[string]string{"name":"john","age":"20","city":"bangalore"}
fmt.Println(maps)
// using make function
code:=make(map[string]int)
fmt.Println(code)
code["python"]=1
code["java"]=2
code["c"]=3
fmt.Println(code)
fmt.Println(code["python"])
//getting a key
value,ok:=code["python"]
fmt.Println(value,ok)
value,ok=code["c++"]
fmt.Println(value,ok)
//adding a value
code["c++"]=4
//override a value
code["python"]=5
//delete a value
delete(code,"c++")

//iterate over all values
for key,val:=range code{
	fmt.Println(key,"=>",val)
}

//truncate a value
code=make(map[string]int)
fmt.Println(code)
}
