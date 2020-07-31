# pointers - summary

1. Pointer type uses "*" as prefix to type it is pointed to. "&" denotes address of memory

2. Dereference a pointer by precending with "*". Complex types are automatically dereference 

3. Pointer arithmatic is not allowed in go (e.g: a := 3 and (b := &a + 3 -> Error)). We can use unsafe library to work with pointer arithmatic

4. Zero value for pointer is nil. If we dont initialize pointer then nil will be there by default

5. Map and Slices are by default having pointer reference. While struct and arrays are copy by value. We externally have to assign pointer type to struct and array