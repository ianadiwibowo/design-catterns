package main

import "fmt"

type Cat struct {
	Name  string
	State State
}

type State int

const (
	Arrived State = iota
	Registered
	WellFed
	Bathed
	FurTrimmed
	Pedicured
	Massaged
)

type CatService interface {
	Handle(cat *Cat)
	SetNext(service CatService)
}

type Reception struct {
	Next CatService
}

func (r *Reception) Handle(cat *Cat) {
	if cat.State != Arrived {
		fmt.Println(cat.Name, "'s state is not arrived")
		return
	}

	fmt.Println("Reception is registering", cat.Name)
	cat.State = Registered

	if r.Next != nil {
		r.Next.Handle(cat)
	}
}

func (r *Reception) SetNext(service CatService) {
	r.Next = service
}

type DiningRoom struct {
	Next CatService
}

func (d *DiningRoom) Handle(cat *Cat) {
	if cat.State != Registered {
		fmt.Println(cat.Name, "'s state is not registered")
		return
	}

	fmt.Println("Dining room is feeding", cat.Name, "well")
	cat.State = WellFed

	if d.Next != nil {
		d.Next.Handle(cat)
	}
}

func (d *DiningRoom) SetNext(service CatService) {
	d.Next = service
}

type Shower struct {
	Next CatService
}

func (s *Shower) Handle(cat *Cat) {
	if cat.State != WellFed {
		fmt.Println(cat.Name, "'s state is not well fed")
		return
	}

	fmt.Println("Shower is bathing", cat.Name)
	cat.State = WellFed

	if s.Next != nil {
		s.Next.Handle(cat)
	}
}

func (s *Shower) SetNext(service CatService) {
	s.Next = service
}

func main() {
	reception := &Reception{}
	diningRoom := &DiningRoom{}
	shower := &Shower{}
	reception.SetNext(diningRoom)
	diningRoom.SetNext(shower)

	cat := &Cat{Name: "Pupuru"}

	reception.Handle(cat)
}
