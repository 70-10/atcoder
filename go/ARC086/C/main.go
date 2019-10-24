package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(logic(a, k))
}

func logic(nums []int, k int) int {
	dic := makeDicionary(nums)
	sort.Sort(dic)

	if dic.Len() <= k {
		return 0
	}

	sum := 0

	for i := 0; i < dic.Len()-k; i++ {
		sum += dic[i].Value
	}
	return sum
}

func makeDicionary(nums []int) Dictionary {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	i := 0
	dic := make(Dictionary, len(m))
	for k, v := range m {
		dic[i] = kv{Key: k, Value: v}
		i++
	}

	return dic
}

type kv struct {
	Key   int
	Value int
}

type Dictionary []kv

func (d Dictionary) Len() int {
	return len(d)
}

func (d Dictionary) Less(i, j int) bool {
	return d[i].Value < d[j].Value
}

func (d Dictionary) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
