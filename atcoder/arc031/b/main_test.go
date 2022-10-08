package main

import (
	"strings"
	"testing"
)

func TestLandfill(t *testing.T) {
	tests := map[string]struct {
		in   string
		size int
		want bool
	}{
		"one_island": {
			in: `xxxxxxxxxx
xoooooooxx
xxoooooxxx
xxxoooxxxx
xxxxoxxxxx
xxxxxxxxxx
xxxxoxxxxx
xxxoooxxxx
xxoooooxxx
xxxxxxxxxx`,
			size: 10,
			want: true,
		},
		"not_one_island": {
			in: `xxxxxxxxxx
xoooooooxx
xxoooooxxx
xxxoooxxxx
xxxxxxxxxx
xxxxxxxxxx
xxxxxxxxxx
xxxoooxxxx
xxoooooxxx
xxxxxxxxxx`,
			size: 10,
			want: false,
		},
		"not_one_island2": {
			in: `xxxxoxxxxx
xxxxoxxxxx
xxxxoxxxxx
xxxxoxxxxx
ooooxooooo
xxxxoxxxxx
xxxxoxxxxx
xxxxoxxxxx
xxxxoxxxxx
xxxxoxxxxx`,
			size: 10,
			want: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := landFill(genA(strings.NewReader(tc.in), tc.size))
			if got != tc.want {
				t.Errorf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}

func TestIsOneIsland(t *testing.T) {
	tests := map[string]struct {
		in   string
		size int
		want bool
	}{
		"one_island": {
			in: `xxxxxxxxxx
xoooooooxx
xxoooooxxx
xxxoooxxxx
xxxxoxxxxx
xxxxoxxxxx
xxxxoxxxxx
xxxoooxxxx
xxoooooxxx
xxxxxxxxxx`,
			size: 10,
			want: true,
		},
		"not_one_island": {
			in: `xxxxxxxxxx
xoooooooxx
xxoooooxxx
xxxoooxxxx
xxxxoxxxxx
xxxxxxxxxx
xxxxoxxxxx
xxxoooxxxx
xxoooooxxx
xxxxxxxxxx`,
			size: 10,
			want: false,
		},
		"not_one_island2": {
			in: `xox
oxx
oxx`,
			size: 3,
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isOneIsland(genA(strings.NewReader(tc.in), tc.size))
			if got != tc.want {
				t.Errorf("got: %t, want: %t", got, tc.want)
			}
		})
	}
}
