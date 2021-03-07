package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	girl := SchoolGirl{"嬌嬌"}
	proxy := NewProxy(girl)
	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()

	// proxy2 := Proxy{Pursuit{SchoolGirl{"123"}}}
	// proxy2.GiveChocolate()
}
