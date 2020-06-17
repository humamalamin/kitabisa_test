package services

import "github.com/kitabisa_test/domains"

type PrimeService struct {
	domains.RowNumber
}

func (ss PrimeService) Perform() {
	ss.Prime()
}
