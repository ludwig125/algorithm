package main

import (
	"reflect"
	"testing"
)

func TestJoinWords(t *testing.T) {
	tests := map[string]struct {
		ss   []string
		want string
	}{
		"FisHDoGCaTAAAaAAbCAC": {
			ss:   []string{"AA", "AaA", "AbC", "AC", "CaT", "DoG", "FisH"},
			want: "AAAaAAbCACCaTDoGFisH",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := joinWords(tc.ss)
			//fmt.Println(got)
			if got != tc.want {
				t.Errorf("got: %v, wants: %v", got, tc.want)
			}
		})
	}
}

func TestSortWords(t *testing.T) {
	tests := map[string]struct {
		ss    []string
		wants []string
	}{
		"FisHDoGCaTAAAaAAbCAC": {
			ss:    []string{"FisH", "DoG", "CaT", "AA", "AaA", "AbC", "AC"},
			wants: []string{"AA", "AaA", "AbC", "AC", "CaT", "DoG", "FisH"},
		},
		"AAAAAjhfgaBCsahdfakGZsZGdEAA": {
			ss:    []string{"AA", "AA", "AjhfgaB", "CsahdfakG", "ZsZ", "GdE", "AA"},
			wants: []string{"AA", "AA", "AA", "AjhfgaB", "CsahdfakG", "GdE", "ZsZ"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortWords(tc.ss)
			//fmt.Println(got)
			if !reflect.DeepEqual(got, tc.wants) {
				t.Errorf("got: %v, wants: %v", got, tc.wants)
			}
		})
	}
}

func TestExtractWords(t *testing.T) {
	tests := map[string]struct {
		s     string
		wants []string
	}{
		"AaA": {
			s:     "AaA",
			wants: []string{"AaA"},
		},
		"AaABaaB": {
			s:     "AaABaaB",
			wants: []string{"AaA", "BaaB"},
		},
		"FisHDoGCaTAAAaAAbCAC": {
			s:     "FisHDoGCaTAAAaAAbCAC",
			wants: []string{"FisH", "DoG", "CaT", "AA", "AaA", "AbC", "AC"},
		},
		"AAAAAjhfgaBCsahdfakGZsZGdEAA": {
			s:     "AAAAAjhfgaBCsahdfakGZsZGdEAA",
			wants: []string{"AA", "AA", "AjhfgaB", "CsahdfakG", "ZsZ", "GdE", "AA"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := extractWords(tc.s)
			if !reflect.DeepEqual(got, tc.wants) {
				t.Errorf("got: %v, wants: %v", got, tc.wants)
			}
		})
	}
}

func TestIsUppercase(t *testing.T) {
	tests := map[string]struct {
		s    rune
		want bool
	}{
		"AisUpper": {
			s:    'A',
			want: true,
		},
		"ZisUpper": {
			s:    'Z',
			want: true,
		},
		"aisLower": {
			s:    'a',
			want: false,
		},
		"zisLower": {
			s:    'z',
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isUppercase(tc.s)
			if got != tc.want {
				t.Errorf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}
