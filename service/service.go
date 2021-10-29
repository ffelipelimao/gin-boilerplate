package service

import "fmt"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) Hey() {
	fmt.Println("Hey")
}
