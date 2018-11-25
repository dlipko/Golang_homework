package main

// сюда вам надо писать функции, которых не хватает, чтобы проходили тесты в gotchas_test.go

import (
	"fmt"
	"sort"
	"strings"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(slice []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", "", -1), "[]")
}

func MergeSlices(sl1 []float32, sl2 []int32) []int {
	var newSl []int
	for _, val := range sl1 {
		newSl = append(newSl, int(val))
	}

	for _, val := range sl2 {
		newSl = append(newSl, int(val))
	}
	return newSl
}

func GetMapValuesSortedByKey(m map[int]string) []string {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var res []string
	for _, k := range keys {
		res = append(res, m[k])
	}

	return res
}
