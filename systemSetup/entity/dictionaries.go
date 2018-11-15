package entity

// 字典类型
type DictionaryType struct {
	Id         int    `db:"id" json:"id"`
	TypeId     int    `db:"type_id" json:"typeId" field:"type_id"`
	Key        string `db:"key" json:"key" field:"key"`
	Title      string `db:"title" json:"title" field:"title"`
	IsUse      bool   `db:"is_use" json:"isUse" field:"is_use"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	Describe   string `db:"describe" json:"describe" field:"describe"`
}

type DelDictionaryType struct {
	Id     int  `db:"id" json:"id"`
	Delete bool `db:"delete" json:"-" field:"delete"`
}

type AddDictionaryType struct {
	TypeId     int    `db:"type_id" json:"typeId" field:"type_id"`
	Key        string `db:"key" json:"key" field:"key"`
	Title      string `db:"title" json:"title" field:"title"`
	IsUse      bool   `db:"is_use" json:"isUse" field:"is_use"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
	UserId     int    `db:"user_id" json:"userId" field:"user_id"`
	Describe   string `db:"describe" json:"describe" field:"describe"`
}

// 字典
type Dictionary struct {
	Id       int    `db:"id" json:"id"`
	TypeId   int    `db:"type_id" json:"typeId" field:"type_id"`
	Key      string `db:"key" json:"key" field:"key"`
	Value    string `db:"value" json:"value" field:"value"`
	Order    int    `db:"order" json:"order" field:"order"`
	Describe string `db:"describe" json:"describe" field:"describe"`
}

type DelDictionary struct {
	Id     int  `db:"id" json:"id"`
	Delete bool `db:"delete" json:"-" field:"delete"`
}

type AddDictionary struct {
	TypeId   int    `db:"type_id" json:"typeId" field:"type_id"`
	Key      string `db:"key" json:"key" field:"key"`
	Value    string `db:"value" json:"value" field:"value"`
	Order    int    `db:"order" json:"order" field:"order"`
	Describe string `db:"describe" json:"describe" field:"describe"`
}
