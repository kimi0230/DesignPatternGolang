/*
	Prototype 原型模式：用原型實例指定建立物件的種類, 並且透過拷貝這些原型建立新的物件
	從一個物件在建立另一個可訂製的物件,而且不需知道任何建立的細節
*/

package prototype

import "fmt"

type WorkExperience struct {
	workDate string
	company  string
}

func (w *WorkExperience) GetWorkDate() string {
	return w.workDate
}

func (w *WorkExperience) SetWorkDate(val string) {
	w.workDate = val
}

func (w *WorkExperience) GetCompany() string {
	return w.company
}

func (w *WorkExperience) SetCompany(val string) {
	w.company = val
}

func (w *WorkExperience) Clone() *WorkExperience {
	if w == nil {
		return nil
	}
	// 淺複製, 複製結構, 但不複製值
	shallowObj := w
	fmt.Printf("&shallowObj: %p, orgObj:  %p , shallowObj: %+v \n", &shallowObj, w, shallowObj)
	return shallowObj
}

type Resume struct {
	name string
	sex  string
	age  string
	work WorkExperience
}

// NewResume 建構值
func NewResume(n string) *Resume {
	work := WorkExperience{}
	return &Resume{name: n, work: work}
}

// 提供clone 方法呼叫的私有建構是, 以便複製工作經歷的資料
func (r *Resume) cloneWork(w WorkExperience) {
	r.work = *w.Clone()
}

// SetPersonalInfo 設定個人資訊
func (r *Resume) SetPersonalInfo(sex, age string) {
	if r == nil {
		return
	}
	r.age = age
	r.sex = sex
}

// SetWorkExperience 設定工作經歷
func (r *Resume) SetWorkExperience(workDate, company string) {
	if r == nil {
		return
	}
	r.work.company = company
	r.work.workDate = workDate
}

// Display 顯示履歷
func (r *Resume) Display() {
	if r == nil {
		return
	}
	fmt.Println(r.name, r.sex, r.age)
	fmt.Println("工作經歷: ", r.work.workDate, r.work.company)
}

// Clone 深複製履歷
// 淺複製: 被複製物件的所有變數都含有與原來的物件相同的值, 而所有的對其物件的參考都能指向
// 深複製: 把參考物件的變數指向複製過的新物件, 而不是原有他被參考的物件
func (r *Resume) Clone() *Resume {
	// 深複製
	// var obj = new(Resume)
	// obj.work = *(r.work).Clone()
	// obj.age = r.age
	// obj.name = r.name
	// obj.sex = r.sex
	// return obj

	// 淺複製 : 會導致來源的值也被更改
	shallowObj := r
	shallowObj.age = r.age
	shallowObj.name = r.name
	shallowObj.sex = r.sex
	shallowObj.work = *(r.work).Clone()
	return shallowObj
}
