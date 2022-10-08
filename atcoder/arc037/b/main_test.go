package main

import (
	"fmt"
	"testing"
)

func TestHoge(t *testing.T) {
	tests := map[string]struct {
		n int
		m int
		g [][]int
		//start int
	}{
		"one_tree": {
			n: 8,
			m: 7,
			g: [][]int{
				[]int{},
				[]int{2},
				[]int{1, 3, 4},
				[]int{2},
				[]int{2},
				[]int{6},
				[]int{5, 7, 8},
				[]int{6, 8},
				[]int{6, 7},
			},
			//start: 5,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			for i, v := range tc.g {
				fmt.Println(i, v)
			}
			treeNum:=0
			checked := make( map[int]bool,8)
			for i:=1;i<=8;i++{
				if checked[i]{
					continue
				}

				ok,reached :=isTree(tc.n,  tc.g,i)
				fmt.Println("got",ok,reached)
				if ok{
					treeNum++
				}
				for i, v := range reached {
					if v==1{
						checked[i]=true
					}
				}
			}
			fmt.Println("treeNum",treeNum)
		})
	}
}

// func TestIsClosed(t *testing.T) {
// 	tests := map[string]struct {
// 		g     [][]int
// 		start int
// 		want  bool
// 	}{
// 		"closed": {
// 			g: [][]int{
// 				[]int{},
// 				[]int{2},
// 				[]int{1, 3, 4},
// 				[]int{2, 4},
// 				[]int{2, 3},
// 			},
// 			want: true,
// 		},
// 		"not_closed": {
// 			g: [][]int{
// 				[]int{},
// 				[]int{2},
// 				[]int{1, 3, 4},
// 				[]int{2},
// 				[]int{2},
// 			},
// 			want: false,
// 		},
// 	}
// 	for name, tc := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			for i, v := range tc.g {
// 				fmt.Println(i, v)
// 			}
// 			fmt.Println("-----")
// 			reached := make([]int, len(tc.g))
// 			got := isClosed(tc.g, 1, &reached, 0)
// 			//			fmt.Println(got)
// 			if got != tc.want {
// 				t.Errorf("got: %t, want:%t", got, tc.want)
// 			}
// 		})
// 	}
// }
