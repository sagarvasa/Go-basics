# variables - Summary

1. Multiple way to declare and use the variable
    1. var start int;
    2. var start int = 0;
    3. start := 0;

    Variable declared at package level can not include ":=" syntax. We need to use full syntax

2. All local variables must be used. Unused variable under function throws compilation error

3. We can shadow the variable but can't re-declare it under same scope

4. Variable visibility
    1. if first letter uppercase -> export
    2. if first letter lowercase -> package scope
    3. no explicit private scope but can use closure concept

5. Scoping
    1. package level + exported
    2. package level only
    3. block level

6. Type conversion is allowed and possible using inbuilt function int(), float32() etc. for string -> need to use strconv package's methods. Lossy conversion is not allowed

