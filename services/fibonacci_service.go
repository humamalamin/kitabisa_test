package services

import "github.com/kitabisa_test/domains"

type FibonacciService struct {
	domains.RowNumber
}

func (ss FibonacciService) Perform() {
	ss.Fibonacci()
}
