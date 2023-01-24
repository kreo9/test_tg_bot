package product

import (
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx >= 0 && idx < len(allProducts) {
		return &allProducts[idx], nil
	}
	return nil, fmt.Errorf("product[%v] not exist", idx)
}
