package models

type A struct {
	Id int `xorm:"integer"`
}

type B struct {
	Id int `xorm:"INTEGER"`
}
