package main

import (
	"fmt"
)

//归并排序
var aux StringArr

type StringArr []string

func (s StringArr) Len() int           { return len(s) }
func (s StringArr) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StringArr) Less(i, j int) bool { return s[i] < s[j] }

func merge(a StringArr, lo, mid, hi int) {
	var (
		i = lo
		j = mid + 1
	)
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

//自顶向下排序
func sortDown(a StringArr, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sortDown(a, lo, mid)
	sortDown(a, mid+1, hi)
	merge(a, lo, mid, hi)
}

func sortUp(a StringArr) {
	N := a.Len()
	for sz := 1; sz < N; sz = sz + sz { //sz 子数组大小
		for lo := 0; lo < N-sz; lo += sz + sz {
			minN := N - 1
			if lo+sz+sz-1 < N-1 {
				minN = lo + sz + sz - 1
			}
			merge(a, lo, lo+sz-1, minN)
		}
	}
}

func main() {
	a := StringArr{"M", "E", "R", "G", "E", "S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	aux = make(StringArr, a.Len())
	//sortDown(a, 0, a.Len()-1)
	sortUp(a)
	fmt.Printf("%v", a)
}
