package main

import "fmt"

// Flyweight objects, immutable, storing intrinsic properties of an object that
// can be shared with other similar objects, thus reducing memory load
// especially for big computations, e.g. big simulation.
type CatAsset struct {
	Sprite []byte // Storing shared expensive animation images of a cat.
}

// The normal object that depends on the flyweight object.
type Cat struct {
	Name string
	*CatAsset
}

// In real life, this will be a serious animation rendering function.
func Render(cats []Cat) {
	for _, c := range cats {
		fmt.Println(c)
	}
}

func main() {
	threeColorCatAsset := &CatAsset{} // Load some shared asset.
	orangeCatAsset := &CatAsset{}     // Load some other different asset.

	// Let's say this is a simulation of 1 billion cats interacting in a city.
	// Instead of each Cat object storing its own animation asset, which is an
	// expensive and memory-intensive operations, they share a flyweight object
	// to store same immutable properties; in this case the animation sprites.
	cats := []Cat{
		{"Kupita", orangeCatAsset},
		{"Lupita", threeColorCatAsset},
		{"Rijong", orangeCatAsset},
	}

	Render(cats) // A serious animation rendering function.
}
