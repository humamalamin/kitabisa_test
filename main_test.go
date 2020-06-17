package main

import (
	"testing"

	"github.com/humamalamin/kitabisa_test/domains"
	"github.com/humamalamin/kitabisa_test/services"
)

var (
	inputSeharusnya [3]int
)

func TestAddition(t *testing.T) {
	calc := domains.Calculator{
		A: 1,
		B: 2,
	}

	jobAddition := services.SumService{calc}

	services.Operator(jobAddition)

	t.Logf("Sum => %d", 3)
}

func TestMultiply(t *testing.T) {
	calc := domains.Calculator{
		A: 1,
		B: 2,
	}

	jobMultiply := services.MultiplyService{calc}

	services.Operator(jobMultiply)

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
