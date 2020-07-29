# array - Summary

1. Array in GO are fixed sized and are of same type

2. They are copy by value by default

3. We can use built in functions like len(), cap() to identify length and capacity of array

4. Multiple ways to use array
    1. var arr = [size]dataType{...values}
    2. arr := [size]dataType{...values}
    3. var arr = [...]dataType{...values}
    4. var arr = [3]dataType and then arr[0], arr[1], arr[2]
    5. Copy array into other array (call by value)
    6. Pass reference of arr to other array (call by reference)