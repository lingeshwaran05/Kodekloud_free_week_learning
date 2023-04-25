echo "exit codes"
if 5 -ls 3
then 
echo "5 is greater than 3"
exit 1
else 
echo "5 is not greater than 3"
exit 0
fi