package main

import (
	"fmt"

	"github.com/humamalamin/kitabisa_test/domains"
	"github.com/humamalamin/kitabisa_test/services"
)

func main() {
	var a, b, c int

	fmt.Println("Enter number : ")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	calc := domains.Calculator{
		A: a,
		B: b,
	}

	jobPenambahan := services.SumService{calc}
	jobPerkalian := services.MultiplyService{calc}

	services.Operator(jobPenambahan)
	services.Operator(jobPerkalian)

	fmt.Println("Enter Range Number: ")
	fmt.Scanf("%d", &c)

	calc_range := domains.RowNumber{
		A: c,
	}

	jobPrime := services.PrimeService{calc_range}
	jobFibonacci := services.FibonacciService{calc_range}

	services.Operator(jobPrime)
	services.Operator(jobFibonacci)
}
