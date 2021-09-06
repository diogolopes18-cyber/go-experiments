package main

import "fmt"

//Declare blocks of variables

var (
	num int = 10
)

var (
	first complex64 = 1 + 2i
)

func main() {
	//The := operator tells the compiler to assert what kind of data type our variable is
	i := 2
	fmt.Println(i)
	diff_format_types()
	test()

	var number int = 10
	fmt.Printf("%T\n", number)
	convert_data_type(number)
	complex_numbers()
}

func diff_format_types() {
	//Go uses the fmt.Printf to format what we want to print, as printf in C
	//%v prints out the value
	//%T prints out the data type

	var j int = 32
	fmt.Printf("%v, %T\n", j, j)
}

//In Go, the innermost variables take precedence, that's called shadowing

func test() {
	//The num variable is declared at package level as 10, but I reassigned it and it takes the value of the innermost variable
	num = 13
	fmt.Printf("%v\n", num)
}

//In Go, if the variables are not assigned a value they are initialiazed with the zero value
func convert_data_type(number int) {
	var new_number float32 = float32(number)
	fmt.Printf("%T\n", new_number)
}

func complex_numbers() {
	//In Go, we can split a complex number into real and complex part
	fmt.Printf("Here's the real part %v\n", real(first))
	fmt.Printf("Here's the complex part %v\n", imag(first))
}
