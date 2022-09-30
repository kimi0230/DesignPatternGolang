package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	var extrinsicState = 22
	fwFactory := NewFlyweightFactory()
	fx := fwFactory.Flyweight("X")
	fx.Operation(extrinsicState)

	fy := fwFactory.Flyweight("Y")
	extrinsicState--
	fy.Operation(extrinsicState)

	fz := fwFactory.Flyweight("Z")
	extrinsicState--
	fz.Operation(extrinsicState)

	uFactory := new(UnsharedConcreteFlyweight)
	extrinsicState--
	uFactory.Operation(extrinsicState)
}
