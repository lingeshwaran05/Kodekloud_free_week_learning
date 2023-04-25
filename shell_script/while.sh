echo "while loop in shell script"
i=1
while 1 -le 5
do
    echo "Welcome $i times"
    i=$((i+1))
done

echo "while with condition and break in shell script"
i=1
while [ $i -le 5 ]
do
    echo "Welcome $i times"
    if [ $i -eq 3 ]
    then
        break
    fi
    i=$((i+1))
done

echo "while with condition and continue "
i=1
while [ $i -le 5 ]
do
    if [ $i -eq 3 ]
    then
        i=$((i+1))
        continue
    fi
    echo "Welcome $i times"
    i=$((i+1))
done
