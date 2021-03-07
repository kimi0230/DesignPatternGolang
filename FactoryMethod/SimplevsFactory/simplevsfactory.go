/*
簡單工廠 vs. 工廠方法
簡單工廠
	最大優點在於工廠類別中包含了必要的邏輯判斷, 根據用戶端的選擇條件動態實體化相關的類別
	對用戶端來說, 去除了與具體產品的依賴
	但是違背了 開放封閉原則
工廠模式
	用戶端需決定實體化哪一個工廠來實現運算類別, 選擇判斷的問題還是存在
	工廠方法把簡單工廠的內部邏輯判斷移到了用戶端程式碼來進行
	如果要加功能, 本來是改工廠類型, 現在是修改用戶端
	是簡單工廠模式的進一步抽象.
	缺點是每加一個產品就需要加一個產品工廠的類別, 增加額外的開發量
*/
package simplevsfactory

import "fmt"

type NigtingaleInterface interface {
	Sweep()
	Wash()
	BuyRice()
}
type Nigtingale struct {
}

func (n *Nigtingale) Sweep() {
	fmt.Println("掃地")
}
func (n *Nigtingale) Wash() {
	fmt.Println("洗衣")
}
func (n *Nigtingale) BuyRice() {
	fmt.Println("買米")
}

type Undergraduate struct {
	Nigtingale
}

type Volunteer struct {
	Nigtingale
}

type SimpleFactory struct {
}

func (s *SimpleFactory) CreateNightingale(factoryType string) NigtingaleInterface {
	var result NigtingaleInterface
	switch factoryType {
	case "學男丁格爾的大學生":
		result = new(Undergraduate)
		return result
	case "社區義工":
		result = new(Volunteer)
		return result
	}
	return nil
}

type IFactory interface {
	CreateNightingale()
}

type UndergraduateFactory struct{}

func (u *UndergraduateFactory) CreateNightingale() NigtingaleInterface {
	return new(Undergraduate)
}

type VolunteerFactory struct{}

func (u *VolunteerFactory) CreateNightingale() NigtingaleInterface {
	return new(Volunteer)
}
