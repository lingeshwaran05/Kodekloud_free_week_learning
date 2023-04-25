echo "functions in shell script"
function add(){
    echo "addition of $1 and $2 is $(($1+$2))"
}
add 5 3

echo "function with return value"
function add(){
    return $(($1+$2))
}
add 5 3

echo "function solution is assigned to another variable"
function add(){
    return $(($1+$2))
}
sum=$(add 5 3)
echo "sum is $sum"