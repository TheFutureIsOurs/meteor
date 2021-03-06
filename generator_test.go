/*
 * @Author: Daiming Liu (xingrufeng)
 */
package meteor

import (
	"fmt"
	"testing"
	"time"
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

func BenchmarkMeteor(b *testing.B) {
	node, err := NewNode(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node.Generate()
	}
}

func BenchmarkTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now().UnixNano()
	}
}
