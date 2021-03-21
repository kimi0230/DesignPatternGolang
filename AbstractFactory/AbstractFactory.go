package abstractfactory

import "fmt"

type User struct {
	id   int
	name string
}

func (u *User) Id() int {
	if u == nil {
		return -1
	}
	return u.id
}

func (u *User) SetId(id int) {
	if u == nil {
		return
	}
	u.id = id
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

type Department struct {
	id   int
	name string
}

func (d *Department) Id() int {
	if d == nil {
		return -1
	}
	return d.id
}
func (d *Department) SetId(id int) {
	if d == nil {
		return
	}
	d.id = id
}
func (d *Department) Name() string {
	if d == nil {
		return ""
	}
	return d.name
}
func (d *Department) SetName(name string) {
	if d == nil {
		return
	}
	d.name = name
}

type IUser interface {
	Insert(*User)
	GetUser(int) *User
}

type SqlServerUser struct {
}

func (s *SqlServerUser) Insert(u *User) {
	if s == nil {
		return
	}
	fmt.Println("在 Sql Server 中給 User insert 一筆紀錄", u)
}

func (s *SqlServerUser) GetUser(id int) (u *User) {
	if s == nil {
		return nil
	}
	u = &User{id, "kk"}
	fmt.Println("在 Sql Server 中根據 id 得到 User 紀錄", *u)
	return u
}

type AccessUser struct {
}

func (a *AccessUser) Insert(u *User) {
	if a == nil {
		return
	}
	fmt.Println("在 Access 中給 Department insert 一筆紀錄", *u)
}

func (a *AccessUser) GetUser(id int) (u *User) {
	if a == nil {
		return nil
	}
	u = &User{id, "kk"}
	fmt.Println("在 Access 中根據 id 得到 User 紀錄", *u)
	return u
}

// 用於用戶端存取 解除與具體資料庫存取的耦合
type IDepartment interface {
	Insert(*Department)
	GetDepartment(int) *Department
}

// 用於存取 SQL server的 Department
type SqlServerDepartment struct {
}

func (s *SqlServerDepartment) Insert(d *Department) {
	if s == nil {
		return
	}
	fmt.Println("在 Sql Server 中給 Department insert 一筆紀錄", *d)
}

func (s *SqlServerDepartment) GetDepartment(id int) (d *Department) {
	if s == nil {
		return nil
	}
	d = &Department{id, "dd"}
	fmt.Println("在 Sql Server 中根據 id 得到 Department 紀錄", *d)
	return d
}

// 用於存取 Access的 Department
type AccessDepartment struct {
}

func (a *AccessDepartment) Insert(d *Department) {
	if a == nil {
		return
	}
	fmt.Println("在 Access 中給 Department insert 一筆紀錄", *d)
}

func (a *AccessDepartment) GetDepartment(id int) (d *Department) {
	if a == nil {
		return nil
	}
	d = &Department{id, "dd"}
	fmt.Println("在 Access 中根據 id 得到 Department 紀錄", *d)
	return d
}

// 定義一個建立存取 User, Department 物件的抽象工廠介面
type IFactory interface {
	CreateUser() IUser
	CreateDepartment() IDepartment
}

// 實現 IFactory 介面
type SqlServerFactory struct{}

func (s *SqlServerFactory) CreateUser() IUser {
	if s == nil {
		return nil
	}
	u := &SqlServerUser{}
	return u

}
func (s *SqlServerFactory) CreateDepartment() IDepartment {
	if s == nil {
		return nil
	}
	d := &SqlServerDepartment{}
	return d
}

// 實現 IFactory 介面
type AccessFactory struct{}

func (a *AccessFactory) CreateUser() IUser {
	if a == nil {
		return nil
	}
	u := &AccessUser{}
	return u

}
func (a *AccessFactory) CreateDepartment() IDepartment {
	if a == nil {
		return nil
	}
	d := &AccessDepartment{}
	return d
}

// 用簡單工廠來改進抽象工廠
type DataAccess struct {
	db string
}

func (d *DataAccess) CreateUser(db string) IUser {
	if d == nil {
		return nil
	}

	var u IUser

	if db == "sqlserver" {
		u = new(SqlServerUser)
	} else if db == "access" {
		u = new(AccessUser)
	}
	return u
}

func (d *DataAccess) CreateDepartment(db string) IDepartment {
	if d == nil {
		return nil
	}

	var u IDepartment

	if db == "sqlserver" {
		u = new(SqlServerDepartment)
	} else if db == "access" {
		u = new(AccessDepartment)
	}
	return u
}
