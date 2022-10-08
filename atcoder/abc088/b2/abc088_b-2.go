// https://qiita.com/drken/items/fd4e5e3630d0f5859067
package main

import (
	"fmt"
	"sort"
	"sync"
	//"time"
)

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	A := 0
	B := 0
	var lock sync.Mutex
	addAlice := func() {
		lock.Lock()
		defer lock.Unlock()
		if len(a) > 0 {
			A += a[0]
			a = a[1:]
		}
	}
	addBob := func() {
		lock.Lock()
		defer lock.Unlock()
		if len(a) > 0 {
			B += a[0]
			a = a[1:]
		}
	}

	startA := make(chan interface{})
	close(startA)
	toA := make(chan int)
	toB := make(chan int)

	terminated := make(chan interface{})
	//toA <- 1
loop:
	for {
		select {
		case <-terminated:
			break loop
		default:
		}

		// wgを使わないとBのgoroutineが起動する前にAが重複して起動する
		//var wg sync.WaitGroup
		//wg.Add(2)
		go func() {
			//	defer wg.Done()
			<-startA
			fmt.Println("A goroutine")
			addAlice()
			fmt.Printf("A: %d %v\n", A, a)
			toB <- 1
			<-toA
			if len(a) == 0 {
				fmt.Println("len a == 0")
				close(terminated)
			}
		}()
		go func() {
			//	defer wg.Done()
			<-toB
			addBob()
			fmt.Printf("B: %d %v\n", B, a)
			toA <- 1
		}()
		//wg.Wait()
		//time.Sleep(1 * time.Second)
		//time.Sleep(1 * time.Millisecond)
	}

	//fmt.Printf("A: %d B: %d\n", A, B)
	fmt.Println(A - B)
}
