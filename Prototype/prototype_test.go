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

	cloneresume := resume.Clone()
	cloneresume.SetPersonalInfo("男", "27")
	cloneresume.SetWorkExperience("2018-2021", "GBA")
	cloneresume.Display()

	// 如果是深複製 resume不會被改到 work, 反之淺複製會改到work
	resume.Display()
	fmt.Printf("%p ,\t %p \n", &resume, &resume.work)
	fmt.Printf("%p ,\t %p \n", &cloneresume, &cloneresume.work)
}
