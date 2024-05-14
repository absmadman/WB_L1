package main

import "fmt"

func intersect(m1, m2 map[int]int) map[int]int {
	m3 := make(map[int]int)
	if len(m2) > len(m1) {
		m1, m2 = m2, m1
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok && v2 == v1 {
			m3[k1] = v1
		}
	}
	return m3
}

func main() {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := map[int]int{2: 2, 3: 3, 4: 4}
	m3 := intersect(m1, m2)
	fmt.Println(m3)
}
