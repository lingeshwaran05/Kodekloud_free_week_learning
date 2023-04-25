echo "switching in shell script"
echo "Enter a number between 1 and 3"
read -p "enter your choice" choice
if [$choice -eq 1]
then
    echo "you have selected 1"
elif [$choice -eq 2]
then
    echo "you have selected 2"
elif [$choice -eq 3]
then
    echo "you have selected 3"
else
    echo "invalid choice"
fi

echo "switching in shell script"

case $choice in 

1) echo "you have selected 1";;
2) echo "you have selected 2";;
3) echo "you have selected 3";;
4) echo break;;
5) echo continue;;
*) echo "invalid choice";;

esac