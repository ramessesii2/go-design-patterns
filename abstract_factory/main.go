package main

import "fmt"

func main() {
	addidasFactory, _ := getSportsFactory("addidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	addidasShoe := addidasFactory.makeShoe()
	addidasShirt := addidasFactory.makeShirt()

	printShirtDetails(nikeShirt)
	printShirtDetails(addidasShirt)

	printShoeDetails(nikeShoe)
	printShoeDetails(addidasShoe)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
