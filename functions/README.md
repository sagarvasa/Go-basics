# functions - Summary

1. Function Signature
    func nameOfFunction() {
        //body
    }

2. Same as variable, function name with first letter upper case is going to be exported through that package.
3. Opening curly brace should be in same line of func keyword else go compiler will throw error

4. parameterized function
    func paramterizedFunction(name dataType, varibleArgs ...dataType) {
        //body
    }
    func paramterizedFunction(name dataType, varibleArgs ...dataType) returnValueDataType {
        //body
        return name
    }
    func paramterizedFunction(name dataType, varibleArgs ...dataType) (returnValueName returnValueDataType) {
        //body
        return
    }
    func paramterizedFunction(name dataType, varibleArgs ...dataType) (dataType, error) {
        //body
        return name,nil
    }
    
    1. Arguments passed is treated as local variable to that function
    2. We can also pass pointer as a args since it improves optimazation (copying all elements will be time consuming)
    3. variable args must be last parameter of function parameters and its received as a slice
    4. In case of return value being address of memory, runtime will automatically upgrade that variable to heap from stack memory

5. anonymous function or IIFE
    func() {
        //body
    }()
    funcExpression := func() {
        //body
    }
    var funcExpression func() = func() {
        //body
    }

    1. We can use IIFE to define variable in private scope. Closure is also applicable in go
    2. in case of function expression, we need to define function first before calling

6. method

    func (name dataType) sayHello() {
        //body
    }
    func (name *dataType) sayHello() {
        //body
    }
    1. its a basically function that is executing in known context. Known context can be anything from int to struct to pointer of int etc
    2. method body has access of underlying dataType