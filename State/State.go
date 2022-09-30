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
