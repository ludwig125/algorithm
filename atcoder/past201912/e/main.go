package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	rs := makeRelation(n)
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < q; i++ {
		sc.Scan()
		s := sc.Text()

		ss := strings.Split(s, " ")
		if ss[0] == "1" {
			rs.follow(toInt(ss[1]), toInt(ss[2]))
		} else if ss[0] == "2" {
			rs.followAllReturn(toInt(ss[1]))
		}
		if ss[0] == "3" {
			rs.followFollow(toInt(ss[1]))
		}
	}
	rs.print()
}

type relation []string
type relations [][]string

func makeRelation(n int) relations {
	rs := make(relations, n)
	for i := 0; i < n; i++ {
		r := make(relation, n)
		for j := 0; j < n; j++ {
			r[j] = "N"
		}
		rs[i] = r
	}
	return rs
}

func (rs *relations) print() {
	for _, r := range *rs {
		for _, v := range r {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func (rs *relations) follow(a, b int) {
	if b-1 != a-1 { // 自分自身でなければfollow
		(*rs)[a-1][b-1] = "Y"
	}
}

func (rs *relations) followAllReturn(a int) {
	for _, v := range rs.followerList(a) {
		rs.follow(a, v)
	}
}

func (rs *relations) followFollow(a int) {
	for _, x := range rs.followList(a) {
		for _, y := range rs.followList(x) {
			rs.follow(a, y)
		}
	}
}

// xがfollowしている人一覧
func (rs *relations) followList(x int) []int {
	var list []int
	for i, v := range (*rs)[x-1] {
		if v == "Y" {
			list = append(list, i+1)
		}
	}
	return list
}

// xをfollowしている人一覧
func (rs *relations) followerList(x int) []int {
	var list []int
	for i, v := range *rs {
		if v[x-1] == "Y" {
			list = append(list, i+1)
		}
	}
	return list
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
