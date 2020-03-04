package main

import (
	"fmt"
)

type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }

func merge(a StringSlice, lo, mid, hi int) {
	var (
		aux = make(StringSlice,hi)
		i   = lo
		j   = mid + 1
	)
	fmt.Println(a, lo, mid,hi)
	return
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}


	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}

}

func sort(a StringSlice, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi - lo)/2
	sort(a, lo, mid)
	sort(a, mid+1, mid)
	merge(a, lo, mid, hi)
}

func main() {
	a := StringSlice{"M","E","R","G","E","S","O","R","T","E","X","A","M","P","L","E"}
	fmt.Printf("%d",a.Len())
	sort(a, 0, a.Len()-1)
}
