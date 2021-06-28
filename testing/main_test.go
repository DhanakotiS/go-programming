package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		inp []int
		out int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 21},
		{[]int{12, 11, 56, 45, 85, 1, 64, 1}, 275},
		{[]int{1, 2, 4, 7, 10, 4, 14, 15, 56}, 113},
	}

	for _, tc := range testCases {
		res := Add(tc.inp)
		if res != tc.out {
			log.Fatalln("FAIL : Expecting: ", tc.out, "Got:", res)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	lis := []int{12, 11, 56, 45, 85, 1, 64, 1}
	for i := 0; i < b.N; i++ {
		Add(lis)
	}

}
