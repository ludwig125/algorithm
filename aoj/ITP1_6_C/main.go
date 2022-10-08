package main

import (
	"fmt"
)

func main() {
	var n, b, f, r, v int
	fmt.Scan(&n)

	p := newProperty()
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		fmt.Scan(&f)
		fmt.Scan(&r)
		fmt.Scan(&v)
		p.register(b, f, r, v)
	}
	p.print()
}

type property struct {
	buildings []building
}

func (p property) print() {
	for i, b := range p.buildings {
		if i != 0 {
			fmt.Println("####################")
		}
		b.print()
	}
}

func (p *property) register(b, f, r, v int) {
	now := p.buildings[b-1].floors[f-1].rooms[r-1]
	p.buildings[b-1].floors[f-1].rooms[r-1] = now + v
}

func newProperty() property {
	var bs []building
	for i := 0; i < 4; i++ {
		bs = append(bs, newBuilding())
	}
	return property{buildings: bs}
}

type building struct {
	floors []floor
}

func (b building) print() {
	for _, f := range b.floors {
		f.print()
	}
}

func newBuilding() building {
	var fs []floor
	for i := 0; i < 3; i++ {
		fs = append(fs, newFloor())
	}
	return building{floors: fs}
}

type floor struct {
	rooms []int
}

func (f floor) print() {
	for _, r := range f.rooms {
		fmt.Printf(" %d", r)
	}
	fmt.Println()
}

func newFloor() floor {
	var rs []int
	for i := 0; i < 10; i++ {
		rs = append(rs, 0)
	}
	return floor{rooms: rs}
}
