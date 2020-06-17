package main

import (
	"fmt"

	"github.com/humamalamin/kitabisa_test/domains"
	"github.com/humamalamin/kitabisa_test/services"
)

func main() {
	var a, b, c int

	fmt.Println("Enter number X,Y: ")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	calc := domains.Calculator{
		A: a,
		B: b,
	}

	jobAddition := services.SumService{calc}
	jobMultiply := services.MultiplyService{calc}

	services.Operator(jobAddition)
	services.Operator(jobMultiply)

	fmt.Println("Enter number to generate N Prime number and Fibonacci: ")
	fmt.Scanf("%d", &c)

	calc_range := domains.RowNumber{
		A: c,
	}

	jobPrime := services.PrimeService{calc_range}
	jobFibonacci := services.FibonacciService{calc_range}

	services.Operator(jobPrime)
	services.Operator(jobFibonacci)
}
