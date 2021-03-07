package simplevsfactory

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	// 簡單工廠
	var sf *SimpleFactory
	studentA := sf.CreateNightingale("學男丁格爾的大學生")
	studentA.BuyRice()
	volunteerA := sf.CreateNightingale("社區義工")
	volunteerA.Wash()

	// 工廠模式
	undergraduateFactory := UndergraduateFactory{} // 換成義工只要改這裡, 不用再去改switch
	studentB := undergraduateFactory.CreateNightingale()
	studentB.BuyRice()

	volunteerFactory := VolunteerFactory{}
	studentC := volunteerFactory.CreateNightingale()
	studentC.Sweep()
}
