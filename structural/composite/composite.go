package composite

import "fmt"

func Run() {

	fmt.Println("How much is the sallary of the soldiers in the Infantary?")
	infantary := NewInfantary()

	soldier1 := NewSoldier("John", "Private", 1500.00)
	soldier2 := NewSoldier("Jane", "Sergeant", 2500.00)
	soldier3 := NewSoldier("Bob", "Lieutenant", 3500.00)

	infantary.AddSoldier(soldier1)
	infantary.AddSoldier(soldier2)
	infantary.AddSoldier(soldier3)

	totalSallary := infantary.GetSallary()
	fmt.Printf("Total sallary of the soldiers in the Infantary: $%.2f\n", totalSallary)
}

type Soldier interface {
	GetSallary() float64
}

type SoldierImpl struct {
	name    string
	rank    string
	sallary float64
}

func NewSoldier(name, rank string, sallary float64) *SoldierImpl {
	return &SoldierImpl{
		name:    name,
		rank:    rank,
		sallary: sallary,
	}
}

func (s *SoldierImpl) GetName() string {
	return s.name
}

func (s *SoldierImpl) GetRank() string {
	return s.rank
}

func (s *SoldierImpl) GetSallary() float64 {
	return s.sallary
}

type Infantary struct {
	Soldiers []Soldier
}

func NewInfantary() *Infantary {
	return &Infantary{
		Soldiers: make([]Soldier, 0),
	}
}

func (i *Infantary) AddSoldier(soldier Soldier) {
	i.Soldiers = append(i.Soldiers, soldier)
}

func (i *Infantary) GetSallary() float64 {
	totalSallary := 0.0
	for _, soldier := range i.Soldiers {
		totalSallary += soldier.GetSallary()
	}
	return totalSallary
}
