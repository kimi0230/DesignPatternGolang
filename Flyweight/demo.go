package flyweight

import "fmt"

type User struct {
	name string
}

func (u *User) Name() string {
	if u == nil {
		return ""
	}
	return u.name
}

func (u *User) SetName(name string) {
	if u == nil {
		return
	}
	u.name = name
}

// Website : 抽象 Flyweight
type Website interface {
	Use(User)
}

type ConcreteWebSite struct {
	name string
}

func (c *ConcreteWebSite) Use(u User) {
	fmt.Println("網站分類: ", c.name+", 用戶: "+u.Name())
}

type WebSiteFactory struct {
	flyweights map[string]Website
}

func NewWebSiteFactory() *WebSiteFactory {
	fwfactory := WebSiteFactory{make(map[string]Website)}
	return &fwfactory
}

func (w *WebSiteFactory) GetWebSiteFactory(name string) Website {
	if w == nil {
		return nil
	}
	if _, ok := w.flyweights[name]; !ok {
		w.flyweights[name] = &ConcreteWebSite{name}
	}
	return w.flyweights[name]
}

func (w *WebSiteFactory) GetWebSiteCount() int {
	return len(w.flyweights)
}
