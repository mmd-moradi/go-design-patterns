package main

type AddidasFactory struct{}

func (n *AddidasFactory) MakeShoe() ShoeProduct {
	return &AddidasShoe{
		Shoe: Shoe{
			logo: "Addidas",
			size: 10,
		},
	}
}

func (n *AddidasFactory) MakeShirt() ShirtProduct {
	return &AddidasShirt{
		Shirt: Shirt{
			logo: "Addidas",
			size: 10,
		},
	}
}
