package main

import "fmt"

func main() {
	// ######### simple interface ##########
	fmt.Println("*****************************")
	var p payment = paytmPayment{}
	status, err := p.validate()
	fmt.Println("Status and error of paytm struct via interface type payment: "+status, err)
	p = phonepePayment{}
	status1, err1 := p.validate()
	fmt.Println("Status and error of phonepe struct via interface type payment: "+status1, err1)
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
