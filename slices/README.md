# slices (Dynamic array) - Summary

1. Slices in GO are variable sized and they are backed by array

2. They are copy by reference by default

3. We can use built in functions like len(), cap() to identify length and capacity of slices

4. Multiple ways to use slice
    1. var slice = []dataType{...values}
    2. slice := []dataType{...values}
    3. Copy slice into other slice (call by reference)
    4. Pass reference of slice to other slice (call by reference)
    5. using built in make function (recommended) (make([]dataType, len, cap)) (e.g: make([]int, 10) will have len & cap 10, make([]int,5,10) will have len = 5 and cap = 10)

5. Number of elements in slices doesn't neccessarly match with size of backing array. Hence we use cap() to increase underlying array size

6. Capacity of slice doubles if underlying array is full

7. We can append function to add element to slice (if underlying arr is small and full then it will copy all element to new array of double capacity, hence slow down the operation)