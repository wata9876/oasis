package model

//User ユーザー情報
type User struct {
	ID       int    `xorm:"pk not null autoincr int" form:"id" json:"id"`
	Name     string `json:"name" 　　　　　　　xorm:"'name'"`
	Age      int    `json:"age" 　　　　　　　xorm:"'age'"`
	Address  string `json:"address" 　　　　　　　xorm:"'not null address'"`
	Password string `json:"password" 　　　　　　　xorm:"'password'"`
}
