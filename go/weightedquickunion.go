package main

import "fmt"

type UF struct {
	Id []int
	Sz []int
}

func NewUF(size int) *UF {
	uf := new(UF)
	uf.Id = make([]int, size)
	fmt.Printf("allocating... \n")
	for i := 0; i < size; i++ {
		uf.Id[i] = i
	}

	PrintUF(uf)

	return uf
}

func Root(uf *UF, root int) int {
	for root != uf.Id[root] {
		root = uf.Id[root]
	}

	return root
}

func Union(uf *UF, p int, q int) *UF {
	rootq := Root(uf, q)
	rootp := Root(uf, p)

	uf.Id[rootp] = rootq

	PrintUF(uf)

	return uf
}

func Find(uf *UF, p int, q int) bool {
	return Root(uf, q) == Root(uf, p)
}

func PrintUF(uf *UF) {
	fmt.Printf("index  ")
	for i := 0; i < len(uf.Id); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	fmt.Printf("value  ")
	for i := 0; i < len(uf.Id); i++ {
		fmt.Printf("%d ", uf.Id[i])
	}
	fmt.Printf("\n")
}

func main() {
	uf := NewUF(10)
	Union(uf, 2, 4)
	Union(uf, 4, 5)
	Union(uf, 9, 8)
	Union(uf, 1, 4)
	Union(uf, 8, 2)
	fmt.Printf("Is %d connected to %d? %b", 1, 2, Find(uf, 1, 2))
	fmt.Printf("Is %d connected to %d? %b", 8, 5, Find(uf, 8, 5))
}
