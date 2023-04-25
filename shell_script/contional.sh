echo "if [condition] then [statements] fi"
echo "if [condition] then [statements] else [statements] fi"
echo "if [condition] then [statements] elif [condition] then [statements] else [statements] fi"
if 20>30
then
echo "20 is greater than 30"
fi

if 20>30
then
echo "20 is greater than 30"
else
echo "20 is less than 30"
fi

if ["string1"="string2"]
then
echo "string1 is equal to string2"
elif ["string1"!="string3"]
then
echo "string1 is not equal to string3"
elif ["5 -gt 3"]
then "5 is greater than 3"
elif [5 -eq 3]
then "5 is equal to 3"
elif [5 -lt 3]
then "5 is less than 3"
elif [5 -ne 3]
then "5 is not equal to 3"
else
echo "5 is not greater than 3"
fi

if [["abcd"=*bc*]]
then
echo "abcd contains bc"
fi

if true && true
then
echo "true and true"
elif true || true
then
echo "true or true"
elif !true
then