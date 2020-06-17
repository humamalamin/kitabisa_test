package services

import "github.com/humamalamin/kitabisa_test/domains"

type MultiplyService struct {
	domains.Calculator
}

func (ms MultiplyService) Perform() {
	ms.Multiply()
}
