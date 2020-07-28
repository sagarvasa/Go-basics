# variables - Summary

1. Multiple way to declare and use the variable
    var start int;
    var start int = 0;
    start := 0;

    Variable declared at package level can not include ":=" syntax. We need to use full syntax

2. All local variables must be used. Unused variable under function throws compilation error

3. We can shadow the variable but can't re-declare it under same scope

4. Variable visibility
    if first letter uppercase -> export
    if first letter lowercase -> package scope
    no explicit private scope but can use closure concept

5. Scoping
    package level + exported
    package level only
    block level

6. Type conversion is allowed and possible using inbuilt function int(), float32() etc. for string -> need to use strconv package's methods. Lossy conversion is not allowed

