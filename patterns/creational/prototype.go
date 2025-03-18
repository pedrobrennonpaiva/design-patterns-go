package creational

import "fmt"

type INPC interface {
	Print(string)
	Clone() INPC
}

type Npc struct {
	Name string
}

func (n *Npc) Print(s string) {
	fmt.Println(s + n.Name)
}

func (n *Npc) Clone() INPC {
	return &Npc{
		Name: n.Name + " Clone",
	}
}
