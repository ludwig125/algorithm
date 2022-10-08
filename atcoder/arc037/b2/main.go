package main

import "fmt"

func main() {
	n, g := genNums()
	fmt.Println(n, g)
	//	fmt.Println(countTree(n,g))
}

func genNums() (int, [][]int) {
	var n, m int
	fmt.Scan(&n, &m)
	g := make([][]int, n+1)
	var u, v int
	for i := 0; i < m; i++ {
		fmt.Scan(&u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	return n, g
}

func dfs(reached *[]int) {

}

// func countTree(n int, g [][]int) int{
// 	treeNum:=0
// 	checked := make( map[int]bool,n)
// 	for i:=1;i<=n;i++{
// 		if checked[i]{
// 			continue
// 		}

// 		ok,reached :=isTree(n,g,i)
// 		//fmt.Println("got",ok,reached)
// 		if ok{
// 			treeNum++
// 		}
// 		for i, v := range reached {
// 			if v==1{
// 				checked[i]=true
// 			}
// 		}
// 	}
// 	return treeNum
// }

// // グラフとスタート地点を与えたら、そのスタート地点から辿れるグループとそのグループが木かどうかを返す関数
// func isTree(n int, g [][]int,start int) (bool,[]int){
// 	reached := make([]int, n+1)
// 	cnt := 0
// 	dfs(g, start, &reached, &cnt)

// 	reachedNum := 0
// 	for _, v := range reached {
// 		if v == 1 {
// 			reachedNum++
// 		}
// 	}

// 	// cnt(辺)の数がreachedNum(頂点)-1でなければ木ではなく閉路
// 	//fmt.Println(start, reachedNum, cnt,reached)
// 	if (reachedNum-1) != int(cnt/2){
// 		return false, reached
// 	}
// 	return true, reached
// }

// func dfs(g [][]int, i int, reached *[]int, cnt *int) {
// 	(*reached)[i] = 1
// 	for _, v := range g[i] {
// 		(*cnt)++
// 		if (*reached)[v] == 0 {
// 			dfs(g, v, reached, cnt)
// 		}
// 	}
// }
