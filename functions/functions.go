package main

import "fmt"

func main() {
	fmt.Println("************* Pointers **************")
	callOtherFunc()

	name := "Sagar"
	callParameterFunc(name)

	v1, v2, v3 := 1, 2, 3
	callParameterFuncWithVariableArgs(name, v1, v2, v3)

	returnName := funcWithReturn(name)
	fmt.Println("funcWithReturn name ", returnName)

	nameAdd := funcWithReturnAddress(name)
	fmt.Println("funcWithReturnAddress name ", *nameAdd)

	var namedReturn = funcWithNamedReturn(name)
	fmt.Println("funcWithNamedReturn name ", namedReturn)

	var m1, m2, m3, m4 float32 = 1.0, 1.0, 2.0, 3.0
	errResult, err := funcWithMultiReturn(m1, m2)
	fmt.Printf("Result -> %v, err -> %v \n", errResult, err)

	result, err1 := funcWithMultiReturn(m3, m4)
	fmt.Printf("Result -> %v, err -> %v \n", result, err1)

	callAnonymousFunc(v1)

	var funcExpression = func() {
		fmt.Println("Function Expression executed")
	}
	funcExpression()

	var funcExpressionWithType func(name string) string
	funcExpressionWithType = func(name string) string {
		returnVal := "Hi"
		fmt.Println("funcExpressionWithType", returnVal, name)
		return returnVal
	}
	funcExpressionWithType(name)

	h1 := greet{
		name: "Sagar",
	}
	h1.sayHello()
	fmt.Println("struct -> ", h1)

	h2 := greet{
		name: "Sagar",
	}
	h2.sayHelloByRef()
	fmt.Println("struct by reference-> ", h2)

	fmt.Println("*************************************")
}

// called from main function
func callOtherFunc() {
	fmt.Println("Other function is called")
}

func callParameterFunc(name string) {
	fmt.Println("callParameterFunc: Hello", name)
}

func callParameterFuncWithVariableArgs(name string, otherArgs ...int) {
	fmt.Printf("callParameterFuncWithVariableArgs Variable type %T and value %v \n", otherArgs, otherArgs)
}

func funcWithReturn(name string) string {
	return name
}

func funcWithReturnAddress(name string) *string {
	returnAdd := name
	return &returnAdd
}

func funcWithNamedReturn(name string) (returnName string) {
	returnName = "Sagar"
	return
}

func funcWithMultiReturn(num1 float32, num2 float32) (float32, error) {
	if num1 == num2 {
		return 0.0, fmt.Errorf("Both values can not be same")
	}
	return num1 / num2, nil
}

func callAnonymousFunc(value int) {
	var m int = 10
	func() {
		fmt.Println("callAnonymousFunc without args", value)
	}()

	func(v int) {
		var j int = 5
		fmt.Println("callAnonymousFunc with args", v)
		fmt.Println("callAnonymousFunc with inner scoping", j, m) //innnerscope have access to outer scope variables
	}(value)

	//fmt.Println("callAnonymousFunc with inner scoping", j) //outerscope doesnt have access to inner scope
}

type greet struct {
	name string
}

func (s greet) sayHello() {
	fmt.Println("Hello", s.name)
	s.name = "monu"
}

func (s *greet) sayHelloByRef() {
	fmt.Println("Hello", s.name)
	s.name = "monu"
}
