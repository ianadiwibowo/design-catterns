package main

import "fmt"

type Cat struct{}

func (c *Cat) EatCatFood(brand string) {
	switch brand {
	case "Whiskas":
		fmt.Println("Yes, eating Whiskas...")
	case "Royal Canin":
		fmt.Println("Yumm, eating Royal Canin delightfully...")
	default:
		fmt.Println("Okay, will eat the cat food later...")
	}
}

// SuperCat decorates (add more functionality) on top of another function
// without changing the original function behavior.
type SuperCat struct {
	*Cat
	FoodRequirementMultiplier int
}

func (s *SuperCat) EatCatFood(brand string) {
	fmt.Println("Wants to eat", s.FoodRequirementMultiplier, "times of cat food...")
	s.Cat.EatCatFood(brand)
	fmt.Println("Supercat asks for more food...")
}

func main() {
	cat := &Cat{}
	cat.EatCatFood("Whiskas") // => Yes, eating Whiskas...

	superCat := &SuperCat{&Cat{}, 3}
	superCat.EatCatFood("Whiskas")
	// => Wants to eat 3 times of cat food...
	// => Yes, eating Whiskas...
	// => Supercat asks for more food...
}
