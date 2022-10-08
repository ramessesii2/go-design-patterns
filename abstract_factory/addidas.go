package main

// concrete factory
type addidas struct {
}

func (a *addidas) makeShoe() iShoe {
	return &addidasShoe{
		shoe: shoe{
			logo: "addidas",
			size: 14,
		},
	}
}

func (a *addidas) makeShirt() iShirt {
	return &addidasShirt{
		shirt: shirt{
			logo: "addidas",
			size: 14,
		},
	}
}
