package main

import (
	"fmt"
	"time"
)

func main() {
	// ######### simple interface ##########
	fmt.Println("*****************************")
	var p payment = paytmPayment{}
	status, err := p.validate()
	fmt.Println("Status and error of paytm struct via interface type payment: "+status, err)
	p = phonepePayment{}
	status1, err1 := p.validate()
	fmt.Println("Status and error of phonepe struct via interface type payment: "+status1, err1)

	// ######### composition interface ##########
	var rw readWrite = newReaderWriterStruct() //constructor function for passing reference on struct methods
	rw.write("String data override")
	r, e := rw.read()
	fmt.Println("composition interface output: "+r, e)

	// ######### own implementation of inbuilt methods ##########
	var writer customWriterIn = customWriterSt{}
	writer.Write([]byte("Hello world"))

	var emptyInf interface{} = 5
	//type emptyInf interface {} //equivalent to above
	switch (emptyInf).(type) {
	case int:
		fmt.Println("integer type")

	default:
		fmt.Println("non integer type")
	}

}

// Payment interface
type payment interface {
	validate() (string, error)
}

// Paytm Struct
type paytmPayment struct {
}

func (paytmPP paytmPayment) validate() (string, error) {
	return "success", nil
}

// Phonepe Struct
type phonepePayment struct {
}

func (phonepe phonepePayment) validate() (string, error) {
	return "failure", nil
}

// ######### composition interface implementation ##########
type reader interface {
	read() (string, error)
}
type writer interface {
	write(string)
}

// interface composition
type readWrite interface {
	reader
	writer
}

type readerWriterStruct struct {
	data *string
}

func (rw *readerWriterStruct) read() (string, error) {
	return *(rw.data), nil //dereferenced
}

func (rw *readerWriterStruct) write(data string) {
	rw.data = &data //writing address of passed data to pointer variable
}

// return pointer to struct
func newReaderWriterStruct() *readerWriterStruct {
	str := "default value set"
	return &readerWriterStruct{
		data: &str,
	}
}

// ######### own implementation of inbuilt methods ##########
type customWriterIn interface {
	Write([]byte)
}

type customWriterSt struct {
}

func (w customWriterSt) Write(bytes []byte) {
	fmt.Printf("[info]: %v:%v:%v - %v \n", time.Now().Hour(), time.Now().Minute(), time.Now().Second(), string(bytes))
}
