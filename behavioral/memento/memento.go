package main

// Cat is the object that will implement the memento function.
// A popular name for memento pattern: audit trails.
type Cat struct {
	Name string
	Age  int
}

// Every time the Cat object is changed, a snapshot will be saved.
func (c *Cat) CreateSnapshot() *CatSnapshot {
	return &CatSnapshot{*c}
}

// In case the Cat object needs to be reverted, it can see its memento list.
func (c *Cat) RestoreSnapshot(snapshot *CatSnapshot) {
	c.Name = snapshot.Copy.Name
	c.Age = snapshot.Copy.Age
}

// This CatSnapshot is a memento (saved state) of the Cat object at
// a certain time.
type CatSnapshot struct {
	Copy Cat
}

// The caretaker takes care of the memento function for the Cat object.
// It maintains the list of CatSnapshots object.
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

	// The initial version of the cat object.
	cat := &Cat{
		Name: "Pupuru",
		Age:  17,
	}
	caretaker.AddSnapshot(cat.CreateSnapshot()) // caretaker size = 1.

	// 2nd version.
	cat.Name = "Pupuruneko"                     // cat => {"Pupuruneko" 17}.
	caretaker.AddSnapshot(cat.CreateSnapshot()) // caretaker size = 2.

	// 3rd version.
	cat.Age = 18                                // cat => {"Pupuruneko" 18}.
	caretaker.AddSnapshot(cat.CreateSnapshot()) // caretaker size = 3.

	// User can revert to the first version, e.g. because of a wrong input.
	cat.RestoreSnapshot(caretaker.GetSnapshot(0)) // caretaker size = 3.
	// Now cat is reverted to => {"Pupuru" 17}.

	// 4th version.
	cat.Age = 19                                // cat => {"Pupuru" 19}.
	caretaker.AddSnapshot(cat.CreateSnapshot()) // caretaker size = 4.
}
