package domains

import "fmt"

type Calculator struct {
	A int
	B int
}

func (c Calculator) Sum() {
	fmt.Printf("%s => %d\n", "Sum", c.A+c.B)
}

func (c Calculator) Multiply() {
	fmt.Printf("%s => %d\n", "Multiply", c.A*c.B)
}

type RowNumber struct {
	A int
}

func (c RowNumber) Prime() {
	var i int
	var number = 2
	count := 0
	var result []int

	for count < c.A {
		div_count := 0
		for i = 1; i <= number; i++ {
			if (number % i) == 0 {
				div_count++
			}
		}

		if div_count < 3 {
			result = append(result, number)
			count = count + 1
		}
		number = number + 1
	}

	fmt.Println("Prime: ", result)
}

func (c RowNumber) Fibonacci() {
	var f1 = 0
	var f2 = 1
	var result []int

	for i := 1; i <= c.A; i++ {
		result = append(result, f1)
		nex := f1 + f2
		f1 = f2
		f2 = nex
	}

	fmt.Println("Fibonacci: ", result)
}
