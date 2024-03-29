### Factory Method 工廠方法模式： 
> 由次類別決定要建立的具象類別為何者
	
定義一個用於創建對象的接口，讓子類決定實例化哪一個類
是一種管理物件創建的模式，隨著輸入的參數不同，簡單工廠會回傳不同的物件
使用者取得物件的時候只要傳入正確的參數，不需要去理解這個物件
是一個使一個類的實例化延遲到其子類

所有應用簡單工廠的地方, 都可以老綠用反射技術來去除 switch 或 if 解除分支判斷帶來的耦合

> Ex: 計算機

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Factory.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/FactoryMethod 

```go
/*
Factory Method 工廠方法模式： 由次類別決定要建立的具象類別為何者
	定義一個用於創建對象的接口，讓子類決定實例化哪一個類
	是一種管理物件創建的模式，隨著輸入的參數不同，簡單工廠會回傳不同的物件
	使用者取得物件的時候只要傳入正確的參數，不需要去理解這個物件
	是一個使一個類的實例化延遲到其子類
此範例為 簡單工廠
*/
package factorymethod

type OperationFunc interface {
	SetNumA(float32)
	SetNumB(float32)
	GetNumA() float32
	GetNumB() float32
	GetResult() (float32, bool)
}

type Operation struct {
	numberA float32
	numberB float32
}

func (o *Operation) SetNumA(num float32) {
	if o == nil {
		return
	}
	o.numberA = num
}

func (o *Operation) SetNumB(num float32) {
	if o == nil {
		return
	}
	o.numberB = num
}
func (o *Operation) GetNumA() float32 {
	if o == nil {
		return 0
	}
	return o.numberA
}
func (o *Operation) GetNumB() float32 {
	if o == nil {
		return 0
	}
	return o.numberB
}

// 加法
type OperationAdd struct {
	Operation
}

func (o *OperationAdd) GetResult() (float32, bool) {
	if o == nil {
		return 0, false
	}
	return o.numberA + o.numberB, true
}

// 減
type OperationSub struct {
	Operation
}

func (o *OperationSub) GetResult() (float32, bool) {
	if o == nil {
		return 0, false
	}
	return o.numberA - o.numberB, true
}

// 乘
type OperationMul struct {
	Operation
}

func (o *OperationMul) GetResult() (float32, bool) {
	if o == nil {
		return 0, false
	}
	return o.numberA * o.numberB, true
}

// 除
type OperationDiv struct {
	Operation
}

func (o *OperationDiv) GetResult() (float32, bool) {
	if o == nil {
		return 0, false
	}
	if o.numberB == 0 {
		return 0, false
	}
	return o.numberA / o.numberB, true
}

// 建立工廠
type OperationFactory struct {
	oper string
}

func (of *OperationFactory) createoperation(op string) OperationFunc {
	switch op {
	case "+":
		return &OperationAdd{}
	case "-":
		return &OperationSub{}
	case "*":
		return &OperationMul{}
	case "/":
		return &OperationDiv{}
	}
	return nil
}

```