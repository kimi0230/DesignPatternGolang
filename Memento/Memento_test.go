package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {

	// 大戰 Boss前
	fmt.Println("大戰 Boss前")
	powerKimi := new(GameRole)
	powerKimi.InitState()
	powerKimi.StateDisplay()

	// 保存進度
	fmt.Println("保存進度")
	stateAdmin := new(RoleStateCaretaker)
	stateAdmin.SetMemento(powerKimi.SaveState())

	// 大戰 Boss, 損耗嚴重
	fmt.Println("大戰 Boss, 損耗嚴重")
	powerKimi.Fight()
	powerKimi.StateDisplay()

	// 恢復之前狀態
	fmt.Println("恢復之前狀態")
	powerKimi.RecoveryState(stateAdmin.Memento())
	powerKimi.StateDisplay()

}
