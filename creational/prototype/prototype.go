package prototype

import "fmt"

type INPC interface {
	Print(string)
	Clone() INPC
}

type Npc struct {
	Name string
}

func Run() {
	npc := Npc{Name: "NPC 1"}
	npc2 := npc.Clone()

	fmt.Println("NPC 1: ", npc.Name)
	npc2.Print("NPC 2: ")
}

func (n *Npc) Print(s string) {
	fmt.Println(s + n.Name)
}

func (n *Npc) Clone() INPC {
	return &Npc{
		Name: n.Name + " Clone",
	}
}
