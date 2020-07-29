# constants - Summary

1. Way to define constant
    const start int = 10;

2. Unlike variable, compiler will not throw error if local constant is unused

3. We can shadow the constant but can't re-declare it under same scope

4. Constant naming - Same as variable
    if first letter uppercase (PascalCase) -> export
    if first letter lowercase (camelCase) -> package scope

5. Constants are replaced at compile time by compiler. Hence value must be calculated at compile time

6. Scoping
    1. package level + exported
    2. package level only
    3. block level

7. various types are
    1. typed constant - (const a int = 5)
    2. untyped constant - (const a = 5) (use compiler ability and can work as a literal)
    3. enumerated constant - (const a = iota)

8. We can not mix two types under same family unless identify using compiler ability. Using compiler ability we can perform implicit conversation unlike variables