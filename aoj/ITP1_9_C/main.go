package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := newScore()
	for i := 0; i < n; i++ {
		var taro, hanako string
		fmt.Scan(&taro, &hanako)
		s.game(taro, hanako)
	}
	s.print()
}

type score struct {
	taro   int
	hanako int
}

func newScore() score {
	return score{taro: 0, hanako: 0}
}
func (s *score) game(t, h string) {
	if t > h {
		s.taro += 3
	} else if t < h {
		s.hanako += 3
	} else {
		s.taro++
		s.hanako++
	}
}

func (s *score) print() {
	fmt.Printf("%d %d\n", s.taro, s.hanako)
}
