package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func (c *Cat) CreateSnapshot() *CatSnapshot {
	return &CatSnapshot{*c}
}

func (c *Cat) RestoreSnapshot(snapshot *CatSnapshot) {
	c.Name = snapshot.Copy.Name
	c.Age = snapshot.Copy.Age
}

type CatSnapshot struct {
	Copy Cat
}

type CatSnapshotCaretaker struct {
	Snapshots []*CatSnapshot
}

func (c *CatSnapshotCaretaker) AddSnapshot(snapshot *CatSnapshot) {
	c.Snapshots = append(c.Snapshots, snapshot)
}

func (c *CatSnapshotCaretaker) GetSnapshot(index int) *CatSnapshot {
	return c.Snapshots[index]
}

func main() {
	caretaker := &CatSnapshotCaretaker{}

	cat := &Cat{
		Name: "Pupuru",
		Age:  17,
	}
	caretaker.AddSnapshot(cat.CreateSnapshot())
	fmt.Println("cat:", cat)
	fmt.Println("caretaker:", caretaker)

	cat.Name = "Pupuruneko"
	caretaker.AddSnapshot(cat.CreateSnapshot())
	fmt.Println("cat:", cat)
	fmt.Println("caretaker:", caretaker)

	cat.Age = 18
	caretaker.AddSnapshot(cat.CreateSnapshot())
	fmt.Println("cat:", cat)
	fmt.Println("caretaker:", caretaker)

	cat.RestoreSnapshot(caretaker.GetSnapshot(0))
	fmt.Println("cat:", cat)
	fmt.Println("caretaker:", caretaker)

	cat.Age = 19
	caretaker.AddSnapshot(cat.CreateSnapshot())
	fmt.Println("cat:", cat)
	fmt.Println("caretaker:", caretaker)
}
