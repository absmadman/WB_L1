package main

import "fmt"

// Работает только с отсортированными массивами
// Возвращает индекс найденного числа
// В случае если числа нет в массиве возвращает -1
func binarySearch(n []int, target int) int {
	l := 0
	r := len(n) - 1
	m := r / 2
	for l <= r {
		if target > n[m] {
			l = m + 1
		} else if target < n[m] {
			r = m - 1
		} else {
			return m
		}
		m = (l + r) / 2
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(a, 5))
}
