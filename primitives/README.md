# primitives - Summary

There are 3 fundamental types
1. Boolean type
2. Numeric type
3. Text type

# Boolean type
    1. either true or false
    2. Default value false if not set
    3. Can not be used as alias for other type (-1 or 0 can not be considered as false)

# Numeric type
1. integer type
    1. Default value is 0
    2. Signed Integer - int8, int16, int32, int64
    3. Unsigned Integer - uint8, uint16, uint32
    4. can not mix two types under same family. (int8 + int16 results in error)
    5. int type --> minimun 32 bits 
    6. uint type --> minimum 32 bits

2. floating type
    1. Default value is 0
    2. different type - float32, float64
    3. can not mix two types under same family. (float32 + float64 results in error)
    4. Literal style - Decimal (1.12), Exponential (1e111), Mixed (1.15e4)
    5. default type under ":=" syntax --> float64

3. complex type
    1. Default value is 0+0i
    2. different type - complex64, complex128
    3. can not mix two types under same family. (complex64 + complex128 results in error)
    4. Built in functions - complex (generate complex no from two float), real (get real part of complex as a float), imag (get imaginary part of complex as a float)
    5. e.g: complex64 -> float 32 real + float 32 imaginary
    6. default type under ":=" syntax --> complex128

# Text type
1. string
    1. UTF8
    2. strings are immutable in go.
    3. many inbuilt method to work with strings
    4. can concat two strings with (+) operator
    5. strings in go are alias for the bytes, hence we can convert to []byte

2. rune
    1. UTF32
    2. Alias for int32
    3. letter := 'a'  (should be in single quotes for rune)