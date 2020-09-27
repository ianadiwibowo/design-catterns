package main

import "fmt"

// A Cat interface has two concrete implementations.
type Cat interface {
	Meow() string
}

type YuenJumboConcreteCat struct{}

func (y *YuenJumboConcreteCat) Meow() string {
	return "A Yuen Jumbo cat is meowing..."
}

type GreenConcreteCat struct{}

func (g *GreenConcreteCat) Meow() string {
	return "A green cat is meowing..."
}

type CatType int

const (
	YuenJumboCat CatType = iota + 1
	GreenCat
)

// This factory method enables the client to defer the decision to get which
// concrete implementation as close to the runtime as possible. As long as the
// return object shares the same interface contract, any type of implementation
// will not break the client. It needs only to depend on the agreed interface.
func NewCat(catType CatType) Cat {
	switch catType {
	case YuenJumboCat:
		return &YuenJumboConcreteCat{}
	case GreenCat:
		return &GreenConcreteCat{}
	default:
		return nil
	}
}

func main() {
	// Whether the Cat object is a YuenJumboConcreteCat or a GreenConcreteCat,
	// the client doesn't need to know nor change its code behavior. And the
	// initialization decision can be delayed as later as possible,
	// like in this place.
	cat1 := NewCat(YuenJumboCat)
	fmt.Println(cat1.Meow())

	cat2 := NewCat(GreenCat)
	fmt.Println(cat2.Meow())
}
