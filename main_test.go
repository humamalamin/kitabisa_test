package main

import (
	"testing"

	"github.com/kitabisa_test/domains"
	"github.com/kitabisa_test/services"
)

var (
	inputSeharusnya [3]int
)

func TestAddition(t *testing.T) {
	calc := domains.Calculator{
		A: 1,
		B: 2,
	}

	jobPenambahan := services.SumService{calc}

	services.Operator(jobPenambahan)

	t.Logf("Sum => %d", 3)
}

func TestMultiply(t *testing.T) {
	calc := domains.Calculator{
		A: 1,
		B: 2,
	}

	jobPerkalian := services.MultiplyService{calc}

	services.Operator(jobPerkalian)

	t.Logf("Muultiply => %d", 2)
}

func TestPrime(t *testing.T) {
	result := [4]int{2, 3, 5, 7}
	calc_range := domains.RowNumber{
		A: 4,
	}

	jobPrime := services.PrimeService{calc_range}

	services.Operator(jobPrime)

	t.Logf("Prime : %d", result)
}

func TestFibonacci(t *testing.T) {
	result := [4]int{0, 1, 1, 2}
	calc_range := domains.RowNumber{
		A: 4,
	}

	jobFibonacci := services.FibonacciService{calc_range}

	services.Operator(jobFibonacci)

	t.Logf("Fibonacci : %d", result)

}
