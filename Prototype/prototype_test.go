package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	resume := NewResume("kimi")

	resume.SetPersonalInfo("男", "24")
	resume.SetWorkExperience("2015-2018", "ABC")
	resume.Display()
	fmt.Println("")

	cloneresume := resume.Clone()
	cloneresume.SetPersonalInfo("女", "18")
	cloneresume.SetWorkExperience("2018-2021", "GBA")
	cloneresume.Display()
	fmt.Println("")

	// 如果是深複製 原本 resume 物件不會被改到
	// 反之淺複製會改到原本的資料
	resume.Display()
	fmt.Println("")
	fmt.Printf("%p ,\t %p \n", &resume, &resume.work)
	fmt.Printf("%p ,\t %p \n", &cloneresume, &cloneresume.work)
}
