package state

import "testing"

func TestState(t *testing.T) {
	work := NewWork()
	work.SetHour(9)
	work.WriteProgram()
	work.SetHour(10)
	work.WriteProgram()
	work.SetHour(12)
	work.WriteProgram()
	work.SetHour(13)
	work.WriteProgram()
	work.SetHour(14)
	work.WriteProgram()
	work.SetHour(17)

	// work.SetFinish(true)
	work.SetFinish(false)
	work.WriteProgram()

	work.SetHour(19)
	work.WriteProgram()
	work.SetHour(22)
	work.WriteProgram()

}
