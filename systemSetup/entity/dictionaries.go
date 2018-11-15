package entity

// 字典类型
type DictionaryType struct {
	AddDictionaryType
	DelDictionaryType
}

type DelDictionaryType struct {
	Id     int32 `db:"id" json:"id" field:"id"`
	Delete bool  `db:"delete" json:"-" field:"delete"`
}

type AddDictionaryType struct {
	TypeId     int32  `db:"type_id" json:"typeId" field:"type_id"`
	Key        string `db:"key" json:"key" field:"key"`
	Title      string `db:"title" json:"title" field:"title"`
	IsUse      bool   `db:"is_use" json:"isUse" field:"is_use"`
	UpdateTime string `db:"update_time" json:"updateTime" field:"update_time"`
	UserId     int32  `db:"user_id" json:"userId" field:"user_id"`
	Describe   string `db:"describe" json:"describe" field:"describe"`
}

// 字典
type Dictionary struct {
	AddDictionary
	DelDictionary
}

type DelDictionary struct {
	Id     int32 `db:"id" json:"id" field:"id"`
	Delete bool  `db:"delete" json:"-" field:"delete"`
}

type AddDictionary struct {
	Key      string `db:"key" json:"key" field:"key"`
	Value    string `db:"value" json:"value" field:"value"`
	Order    int32  `db:"order" json:"order" field:"order"`
	Describe string `db:"describe" json:"describe" field:"describe"`
}
