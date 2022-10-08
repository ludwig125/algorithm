package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	a := genA(os.Stdin, 10)
	if landFill(a) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

// 入力元と縦横の長さを指定する
func genA(r io.Reader, n int) [][]string {
	var a [][]string

	s := bufio.NewScanner(r)
	for i := 0; i < n; i++ {
		s.Scan()
		var b []string
		l := s.Text()
		for j := 0; j < n; j++ {
			b = append(b, string(l[j]))
		}
		a = append(a, b)
	}
	return a
}

func landFill(a [][]string) bool {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// aのコピーを用意
			b := copyFromA(a)
			if b[i][j] == "x" {
				b[i][j] = "o"
			}
			// 1ます埋め立てて全体が一つの島になっていたらTrue
			if isOneIsland(b) {
				return true
			}
		}
	}
	return false
}

// Sliceの値コピーをする関数
func copyFromA(a [][]string) [][]string {
	var b [][]string
	for _, v := range a {
		var l []string
		for _, v2 := range v {
			l = append(l, v2)
		}
		b = append(b, l)
	}
	return b
}

// 一つの島か判定する関数
func isOneIsland(a [][]string) bool {
	n := len(a)

	// aのコピーを用意
	reached := copyFromA(a)
	startX, startY := getStartO(a)

	dfs(n, &reached, startX, startY)
	for _, v := range reached {
		for _, v2 := range v {
			if v2 == "o" {
				return false //oが残っていたら全部塗りつぶせていない
			}
		}
	}
	return true
}

// 最初に見つかった"o"をスタート地点として返す
func getStartO(a [][]string) (int, int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == "o" {
				return j, i
			}
		}
	}
	return 0, 0
}

// startのx,yからはじめて全て"x"にしていく関数
func dfs(n int, r *[][]string, x, y int) {
	if x < 0 || y < 0 || x >= n || y >= n || (*r)[y][x] == "x" {
		return
	}

	if (*r)[y][x] == "x" { // 既に"x"なら戻る
		return
	}
	(*r)[y][x] = "x"
	dfs(n, r, x+1, y)
	dfs(n, r, x-1, y)
	dfs(n, r, x, y+1)
	dfs(n, r, x, y-1)
}
