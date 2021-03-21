package abstractfactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	uData := User{1, "u"}
	dData := Department{1, "d"}
	// 抽象工廠
	fmt.Println("---------- 抽象工廠")
	// factory := new(SqlServerFactory)
	factory := new(AccessFactory)
	iu := factory.CreateUser()
	iu.Insert(&uData)
	iu.GetUser(1)

	// 簡單工廠+抽象工廠
	fmt.Println("---------- 簡單工廠+抽象工廠")
	data := DataAccess{"access"}

	iU := data.CreateUser()
	iU.Insert(&uData)
	gU := iU.GetUser(1)
	fmt.Println(gU)

	data2 := DataAccess{"sqlserver"}
	iD := data2.CreateDepartment()
	iD.Insert(&dData)
	gD := iD.GetDepartment(1)
	fmt.Println(gD)

}
