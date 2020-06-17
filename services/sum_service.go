package services

import "github.com/kitabisa_test/domains"

type SumService struct {
	domains.Calculator
}

func (ss SumService) Perform() {
	ss.Sum()
}
