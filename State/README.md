### State 狀態模式:
> 讓物件的外顯行為隨內部狀態的改變而改變, 彷彿連類別也變了似的

當一個物件的內在狀態改變時, 允許改變其行為,
這個物件看起來像是改變了其類別

主要解決的是當控制一個物件狀態轉換的條件運算式過於複雜的情況.
把狀態的判斷邏輯轉移到表示不同狀態的一系列類別中,
可以把複雜的判斷邏輯簡化

State 的 UML 與策略模式相似
* [Strategy 策略模式：](https://github.com/kimi0230/DesignPatternGolang/tree/master/Strategy): 
    是用在對多個做同樣事情（統一接口）的類對象的選擇上
* [State 狀態模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/State)
    將對某個事情的處理過程抽象成接口和實現類的形式，
    由context保存一份state，在state實現類處理事情時，修改狀態傳遞給context， 由context繼續傳遞到下一個狀態處理中，

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/State.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/State 

```go
package state

import "fmt"

// 工作類別
type Work struct {
	current State
	hour    int
	finish  bool
}

func NewWork() *Work {
	return &Work{current: new(ForenoonState)}
}

func (w *Work) Hour() int {
	if w == nil {
		return -1
	}
	return w.hour
}

func (w *Work) State() State {
	if w == nil {
		return nil
	}
	return w.current
}

func (w *Work) Finish() bool {
	if w == nil {
		return false
	}
	return w.finish
}

func (w *Work) SetHour(h int) {
	if w == nil {
		return
	}
	w.hour = h
}

func (w *Work) SetState(s State) {
	if w == nil {
		return
	}
	w.current = s
}

func (w *Work) SetFinish(f bool) {
	if w == nil {
		return
	}
	w.finish = f
}

func (w *Work) WriteProgram() {
	if w == nil {
		return
	}
	w.current.WriteProgram(w)
}

type State interface {
	WriteProgram(*Work)
}

type ForenoonState struct{}

func (f *ForenoonState) WriteProgram(w *Work) {
	if w.Hour() < 12 {
		fmt.Printf("當前時間 %d 點, 上午工作, 精神百倍 \n", w.Hour())
	} else {
		w.SetState(new(NoonState))
		w.WriteProgram()
	}
}

type NoonState struct{}

func (f *NoonState) WriteProgram(w *Work) {
	if w.Hour() < 13 {
		fmt.Printf("當前時間 %d 點, 餓了, 午飯; 覺得睏, 午休 \n", w.Hour())
	} else {
		w.SetState(new(AfternoonState))
		w.WriteProgram()
	}
}

type AfternoonState struct{}

func (f *AfternoonState) WriteProgram(w *Work) {
	if w.Hour() < 17 {
		fmt.Printf("當前時間 %d 點, 下午狀態還不錯, 繼續努力 \n", w.Hour())
	} else {
		w.SetState(new(EveningState))
		w.WriteProgram()
	}
}

type EveningState struct{}

func (f *EveningState) WriteProgram(w *Work) {
	if w.Finish() {
		// 完成任務轉入下班狀態
		w.SetState(new(RestState))
		w.WriteProgram()
	} else {
		if w.Hour() < 21 {
			fmt.Printf("當前時間 %d 點, 加班喔, 疲累至極 \n", w.Hour())
		} else {
			w.SetState(new(SleepingState))
			w.WriteProgram()
		}
	}
}

type SleepingState struct{}

func (f *SleepingState) WriteProgram(w *Work) {
	fmt.Printf("當前時間 %d 點, 不行了, 睡著了 \n", w.Hour())
}

type RestState struct{}

func (f *RestState) WriteProgram(w *Work) {
	fmt.Printf("當前時間 %d 點, 下班回家了 \n", w.Hour())
}
```