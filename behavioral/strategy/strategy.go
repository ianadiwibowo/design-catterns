package main

import "fmt"

// First, the strategy should be an interface, so the implementation can vary.
type Catcher interface {
	Catch(animal string)
}

// Strategy 1.
type ThreeColorsCatStrategy struct{}

func (t *ThreeColorsCatStrategy) Catch(animal string) {
	fmt.Println("Use telepathic stare to freeze", animal, "then catch it...")
}

// Strategy 2. Both are using the same signature, but different implementations.
type OrangeCatStrategy struct{}

func (o *OrangeCatStrategy) Catch(animal string) {
	fmt.Println("Just pursue bravely the", animal, "until being caught...")
}

// A Cat object doesn't need to implement its own Prey/Catch method but
// can be "injected" with different strategies during runtime; like a plug and
// play weapon. Yes, Strategy Pattern helps to enforce the Dependency Inversion
// Principle through dependency injection.
type Cat struct {
	Skill Catcher
}

// The Prey execution will follow whatever the strategy implements.
// At this point in the definition, a Cat needs to know only the strategy
// interface.
func (c *Cat) Prey(animal string) {
	c.Skill.Catch(animal)
}

// During Cat creation, we inject with different strategy/weapon so
// two Cat objects can have different behaviors.
func main() {
	threeColorsCat := &Cat{&ThreeColorsCatStrategy{}}
	threeColorsCat.Prey("mouse")

	orangeCat := &Cat{&OrangeCatStrategy{}}
	orangeCat.Prey("mouse")
}
