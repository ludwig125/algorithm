package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(joinWords(sortWords(extractWords(s))))
}

func joinWords(ss []string) string {
	return strings.Join(ss, "")
}

func sortWords(ss []string) []string {
	m := make(map[string]string)
	var l []string
	for _, s := range ss {
		m[strings.ToLower(s)] = s
		l = append(l, strings.ToLower(s))
	}
	sort.Strings(l)

	var ret []string
	for _, v := range l {
		ret = append(ret, m[v])
	}
	return ret
}

func extractWords(s string) []string {
	var words []string
	var begin, end int
	cnt := 0
	for i := 0; i < len(s); i++ {
		if isUppercase(rune(s[i])) {
			if cnt == 0 { // 単語先頭
				begin = i
				cnt++
			} else { // 単語末尾
				end = i
				cnt = 0 //cnt reset
				//fmt.Println(begin, end)
				words = append(words, s[begin:end+1])
			}
		}
	}
	return words
}

func isUppercase(r rune) bool {
	if int(r) >= 65 && int(r) <= 90 {
		return true
	}
	return false
}
