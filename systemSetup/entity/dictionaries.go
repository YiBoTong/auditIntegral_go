package entity

import "github.com/golang/protobuf/ptypes/timestamp"

// 字典类型
type DictionaryType struct {
	Id         int32               `db:"id" json:"id"`
	typeId     int32               `db:"type_id" json:"typeId"`
	Key        string              `db:"key" json:"key"`
	Title      string              `db:"title" json:"title"`
	isUse      bool                `db:"is_use" json:"isUse"`
	updateTime timestamp.Timestamp `db:"update_time" json:"updateTime"`
	userId     int32               `db:"user_id" json:"userId"`
	describe   string              `db:"describe" json:"describe"`
}

// 字典
type Dictionary struct {
	Id       int32  `db:"id" json:"id"`
	typeId   int32  `db:"type_id" json:"typeId"`
	Key      string `db:"key" json:"key"`
	Value    string `db:"value" json:"value"`
	Order    int32  `db:"order" json:"order"`
	Describe string `db:"describe" json:"describe"`
}
