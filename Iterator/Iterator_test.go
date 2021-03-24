package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	a := NewConcreteAggregate()
	a.SetThis(0, Object{"Kimi"})
	a.SetThis(1, Object{"Yellow"})
	a.SetThis(2, Object{"International staff"})
	a.SetThis(3, Object{"American"})
	a.SetThis(4, Object{"Package"})

	i := NewConcreteIterator(a)
	// firstItem := i.First()
	// fmt.Println(firstItem.(Object))
	for !i.IsDone() {
		fmt.Println(i.CurrentItem().(Object).Name, " 請買車票！")
		i.Next()
	}
}
