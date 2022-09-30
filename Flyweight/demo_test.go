package flyweight

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	f := NewWebSiteFactory()

	fproduct1 := f.GetWebSiteFactory("產品展示")
	user1 := new(User)
	user1.SetName("Kimi")
	fproduct1.Use(*user1)

	fproduct2 := f.GetWebSiteFactory("產品展示")
	user2 := new(User)
	user2.SetName("YY")
	fproduct2.Use(*user2)

	fblog1 := f.GetWebSiteFactory("部落格")
	user3 := new(User)
	user3.SetName("CC")
	fblog1.Use(*user3)

	fblog2 := f.GetWebSiteFactory("部落格")
	user4 := new(User)
	user4.SetName("DD")
	fblog2.Use(*user4)

	// 給了四個不同用戶使用網站, 實際上只有兩個實體網站
	fmt.Println("得到網站分類總數為: ", f.GetWebSiteCount())
}
