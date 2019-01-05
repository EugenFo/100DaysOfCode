package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i                               // & generates a pointer to its operand (p points to i)
	fmt.Println("Prints value of i:", *p) // prints value of p, which is a pointer to i
	fmt.Println("Prints memory address of i:", p)
	fmt.Println("prints &p:", &p) // i guess its the memory address of p
	*p = 21                       // set the value of p to 21, which will rewrite the value of i, because p points to i
	fmt.Println("Print again memory address of i after set new value:", p)
	fmt.Println(i) // prints value of i, which is 21

	p = &j // p points to j
	fmt.Println("&p after p is pointer of j:", &p)
	fmt.Println("Prints value of j:", *p)
	fmt.Println("Prints memory address of j", p) // print memory address of j, cause p points to j
	*p = *p / 37
	fmt.Println("value of j:", j)
	fmt.Println("values of p:", *p) // prints the actual value of the pointer p, p points to j
	fmt.Println("memory address of j:", p)
	fmt.Println("memory address of p:", &p) // prints the actual memory address of p, p => j

	// changes only the value, not the memory address
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
