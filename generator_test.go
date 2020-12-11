package genuid

import (
	"fmt"
	"testing"
)

func TestSeq(t *testing.T) {
	fmt.Println(maxSecond)
	node, err := NewNode(0)
	fmt.Println(node)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 20; i++ {
		id, _ := node.Generate()
		fmt.Println(id)
	}

}

func TestRand(t *testing.T) {
	var seed uint8 = 1
	for i := 0; i < 10; i++ {
		seed = rand(seed)
		fmt.Println(seed)
	}
	fmt.Printf("%b\t%b\n", 270369, 270369 | 7)
}

func TestRand1(b *testing.T) {
	x := 0
	y := 0
	z := 0
	a := 1
	for i:=0;i<10;i++ {
		t := x ^ (x << 4)
		x = y
		y = z
		z = a
		a = z ^ t ^ ( z >> 1) ^ (t << 1)
		fmt.Println(a)
	}
}