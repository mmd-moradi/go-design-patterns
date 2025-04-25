package main

type NikeFactory struct{}

func (n *NikeFactory) MakeShoe() ShoeProduct {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "Nike",
			size: 10,
		},
	}
}

func (n *NikeFactory) MakeShirt() ShirtProduct {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "Nike",
			size: 10,
		},
	}
}
