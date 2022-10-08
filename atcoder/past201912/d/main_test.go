package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		samples     int
		replaceFrom int
		replaceTo   int
		want        string
	}{
		"1to1": {
			samples:     10,
			replaceFrom: 1,
			replaceTo:   1,
			want:        "Correct",
		},
		"1to2": {
			samples:     10,
			replaceFrom: 1,
			replaceTo:   2,
			want:        "2 1",
		},
		"10to1": {
			samples:     10,
			replaceFrom: 10,
			replaceTo:   1,
			want:        "1 10",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			nums := makeNums(tc.samples)
			nums = replaceAtoB(nums, tc.replaceFrom, tc.replaceTo)
			t.Log(nums)
			got := run(shuffle(nums))
			//t.Logf("got: %s, want: %s", got, tc.want)
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}

func makeNums(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = i + 1
	}
	return ret
}

func replaceAtoB(nums []int, a, b int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == a {
			ret[i] = b
			continue
		}
		ret[i] = nums[i]
	}
	return ret
}

func shuffle(nums []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	return nums
}
