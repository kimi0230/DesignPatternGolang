package builder

import (
	"testing"
)

// |								|     implementing interface By			|
// |								| value receiver | pointer receiver		|
// |constructor returns values		|	ok			 | error				|
// |constructor returns pointers	|	ok			 | ok					|

func TestBuilder(t *testing.T) {
	cbuild1 := new(ConcreteBuilder1)
	director := Director{cbuild1}
	director.Construct()
	director.b.GetResult().Show()

	cbuild2 := new(ConcreteBuilder2)
	director2 := Director{cbuild2}
	director2.Construct()
	director2.b.GetResult().Show()

}
