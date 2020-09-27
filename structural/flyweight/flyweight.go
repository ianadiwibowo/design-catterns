package main

// Flyweight object, immutable, store intrinsic properties of object that
// can be shared with other similar object, thus reducing memory load
// especially for big computations, e.g. simulation
type CatAsset struct {
	Sprite []byte // Storing expensive the animation images of a cat
}

// Normal object that depends on the flyweight object
type Cat struct {
	Name string
	*CatAsset
}

func Render(cats []*Cat) {
	for _, c := range cats {
		fmt.Println(c)
	}
}

func main() {
	threeColorCatAsset := &CatAsset{} // Load some asset
	orangeCatAsset := &CatAsset{}     // Load some different asset

	cats := []&Cat{
		{"Kupita", orangeCatAsset},
		{"Lupita", threeColorCatAsset},
		{"Rijong", orangeCatAsset},
	}

	Render(cats)
}
