/*
Strategy 策略模式： 講可以互換的行為封裝起來, 並使用專接方式決定要使用何者
用戶端實體化的事 CashContext 的物件, 調用的是 CashContext 的方法 GetResult
使具體的收費演算法測底與用戶端分離
*/

package strategy

import "errors"

type CashSuper interface {
	acceptCash(memory float32) float32
}

// CashNomal : 正常收費
type CashNomal struct {
}

func (this *CashNomal) acceptCash(money float32) float32 {
	return money
}

// CashRebate : 打折收費
// meneyRebate 折扣率
type CashRebate struct {
	meneyRebate float32
}

func (this *CashRebate) acceptCash(money float32) float32 {
	return this.meneyRebate * money
}

// CashReturn : 紅利收費
// 比如滿300送100, meneyCondition=300, meneyReturn=100
type CashReturn struct {
	meneyCondition float32
	meneyReturn    float32
}

func (this *CashReturn) acceptCash(money float32) float32 {
	if money >= this.meneyCondition {
		return money - float32(int(money/this.meneyCondition))*this.meneyReturn
	} else {
		return money
	}
}

type CashContext struct {
	CashSuper
}

func (cc *CashContext) CashContext(str string) error {
	switch str {
	case "正常收費":
		cc.CashSuper = &CashNomal{}
	case "滿300送100":
		cc.CashSuper = &CashReturn{300, 100}
	case "打8折":
		cc.CashSuper = &CashRebate{0.8}
	default:
		err := errors.New("no match")
		return err
	}
	return nil
}

func (cc *CashContext) GetResult(money float32) float32 {
	return cc.acceptCash(money)
}
