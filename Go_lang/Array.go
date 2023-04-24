package main
import "fmt"
func main(){
	//Array to store similar type of objects
	var a[5] int
	//var <array name> [size] <type>
	fmt.Println(a)
	a= [5]int{1,2,3,4,5}
	fmt.Println(a)
	var fruits [5]string=[5]string{"apple","banana","mango","orange","grapes"}
	fmt.Println(fruits)
	//using short hand
	//<array name> := [size] <type>{<elements>}
	animals:= [5]string{"dog","cat","cow","horse","goat"}
	fmt.Println(animals)
	names:=[...]string{"john","jane","joe","jim","jerry"}
	fmt.Println(names)
	fmt.Println(len(names))
	for i:=0;i<len(names);i++{
		fmt.Println(names[i])
	}
	//using range
	for i,name:=range names{
		fmt.Println(i,"=>",name)
	}
	//multi dimensional array
	var b[2][3] int
	b= [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(b)
	//using short hand
	c:=[2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(c)

	//slice
	//pointer -> head
	//length -> no of elements len()
	//capacity -> no of elements from head to end cap()
	//<slice name> := []{<type>,<length>,<capacity>}
	//array[start:end]
	slice:=names[0:3]
	fmt.Println(slice)
	//using make function
	fmt.Print(len(names))
	fmt.Print(cap(names))
	sub_slice:=make([]string,3,5)
	fmt.Println(sub_slice)
	//modifying slice
	sub_slice=append(sub_slice,"john","jane","joe")
	fmt.Print(sub_slice)
	sub_slice[3]="jim"
	fmt.Println(sub_slice)
	fmt.Println(len(sub_slice))
	fmt.Println(cap(sub_slice))
	arr:=[5]int{1,2,3,4,5}
	slice1:=arr[:2]
	arr2:=[5]int{6,7,8,9,10}
	slice2:=arr2[2:]
	//using ...
	new_slice:=append(slice1,slice2...)
	fmt.Println(new_slice)
	//delete slice
	new_slice=append(new_slice[:2],new_slice[3:]...)
	fmt.Println(new_slice)
	//copy from a slice
	func copy(des,src,[]int)int{
		copy:=copy(src,des)t
		fmt.Println(copy)
		fmt.Println("number of elements in slice",len(copy))
	}
	des:=[4]int{1,2,3,4}
	src:=[3]int{5,6,7}
	copy(des,src)

}