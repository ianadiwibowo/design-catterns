package main

import "fmt"

type Category int

const (
	Pupuru Category = iota
	Orange
	ThreeColors
	YuenJumbo
	Penguin
	Green
)

// CatBuilder is meant to help simplify the Cat object creation to
// prevent a big and long initialization.
type CatBuilder struct {
	FirstName string
	LastName  string
	Category  Category
}

func NewCatBuilder() *CatBuilder {
	return &CatBuilder{}
}

func (c *CatBuilder) SetFirstName(name string) *CatBuilder {
	c.FirstName = name
	return c
}

func (c *CatBuilder) SetLastName(name string) *CatBuilder {
	c.LastName = name
	return c
}

func (c *CatBuilder) SetCategory(category Category) *CatBuilder {
	c.Category = category
	return c
}

func (c *CatBuilder) Build() *Cat {
	return &Cat{
		Name:     fmt.Sprintf("%s %s", c.FirstName, c.LastName),
		Category: c.Category,
	}
}

// Assume Cat object is a very big and complex object, with a lot of fields.
// The Cat object creation is a big and long initialization.
type Cat struct {
	Name     string
	Category Category
}

func main() {
	builder := NewCatBuilder()
	builder.SetFirstName("Lupita").SetLastName("Neko").SetCategory(ThreeColors)

	cat := builder.Build()
	fmt.Println("Result:", cat) // => Result: &{Lupita Neko 2}
}
