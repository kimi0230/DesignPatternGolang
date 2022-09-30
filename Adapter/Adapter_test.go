package adapter

import "testing"

func TestAdapter(t *testing.T) {
	playA := NewGuards("宮城")
	playA.Attack()

	playB := NewForwards("流川")
	playB.Attack()

	playC := NewTranslator("大猩猩")
	playC.Attack()
}
