package main 
import "fmt"

type UF struct {
	Id []int
}

func NewUF(size int) *UF {
	uf := new(UF)
	uf.Id = make([]int, size)
	fmt.Printf("allocating...\n")
	for i := 0; i < size; i++ {
		uf.Id[i] = i;
	}

	PrintUF(uf)

	return uf
}

func Find(uf *UF, p int, q int) bool {
	fmt.Printf("are %d and %d connected?\n", p, q)

	idp := uf.Id[p]
	idq := uf.Id[q]

	if idp == idq {
		fmt.Printf("yes\n")
		return true
	} else { 
		fmt.Printf("no\n")
		return false
	} 
}

func Union(uf *UF, p int, q int) *UF {
	fmt.Printf("uniting... %d to %d\n", p, q)

	idp := uf.Id[p]
	idq := uf.Id[q]

	for i := 0; i < len(uf.Id); i++ {
		if uf.Id[i] == idp {
			uf.Id[i] = idq
		}
	}

	PrintUF(uf)

	return uf
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
	Union(uf, 1, 5)
	Union(uf, 6, 4)
	Union(uf, 2, 3)
	Union(uf, 6, 5)
	Find(uf, 5, 3)
	Find(uf, 1, 5)
}