package main

type ShoeProduct interface {
	GetLogo() string
	GetSize() int
	SetLogo(logoName string)
	SetSize(size int)
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetSize() int {
	return s.size
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}
