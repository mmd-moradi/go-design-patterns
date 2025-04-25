package main

type ShirtProduct interface {
	GetLogo() string
	GetSize() int

	SetLogo(logoName string)
	SetSize(size int)
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetSize() int {
	return s.size
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}
