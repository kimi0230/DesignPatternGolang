package visitor

import "fmt"

type Action interface {
	GetManConclusion(Man)
	GetWomanConclusion(Woman)
}

type Success struct {
	Name string
}

func (s *Success) GetManConclusion(m Man) {
	if s == nil {
		return
	}
	fmt.Printf("%s %s 時, 背後多半有一個偉大的女人 \n", m.Name, s.Name)
}

func (s *Success) GetWomanConclusion(w Woman) {
	if s == nil {
		return
	}
	fmt.Printf("%s %s 時, 背後多半有一個不成功的男人 \n", w.Name, s.Name)
}

type Failing struct {
	Name string
}

func (s *Failing) GetManConclusion(m Man) {
	if s == nil {
		return
	}
	fmt.Printf("%s %s 時, 悶頭喝酒, 誰也不用勸 \n", m.Name, s.Name)
}

func (s *Failing) GetWomanConclusion(w Woman) {
	if s == nil {
		return
	}
	fmt.Printf("%s %s 時, 眼淚汪汪, 誰也勸不了 \n", w.Name, s.Name)
}

type Person interface {
	Accept(Action)
}

type Man struct {
	Name string
}

func (m *Man) Accept(visitor Action) {
	if m == nil {
		return
	}
	visitor.GetManConclusion(*m)
}

type Woman struct {
	Name string
}

func (w *Woman) Accept(visitor Action) {
	if w == nil {
		return
	}
	visitor.GetWomanConclusion(*w)
}

type ObjectStructurePerson struct {
	elemetes []Person
}

func (o *ObjectStructurePerson) Attach(ele Person) {
	if o == nil || ele == nil {
		return
	}
	o.elemetes = append(o.elemetes, ele)
}

func (o *ObjectStructurePerson) Detach(ele Person) {
	if o == nil || ele == nil {
		return
	}
	for i, val := range o.elemetes {
		if val == ele {
			o.elemetes = append(o.elemetes[:i], o.elemetes[i+1:]...)
			break
		}
	}
}

func (o *ObjectStructurePerson) Accept(v Action) {
	if o == nil {
		return
	}
	for _, val := range o.elemetes {
		val.Accept(v)
	}
}
