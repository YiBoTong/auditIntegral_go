package db

import (
	"auditIntegral/_goSql"
	"auditIntegral/_public/db"
	"auditIntegral/systemSetup/entity"
)

type Dictionaries struct {
	db *gosql.DbMysql
}

func NewDictionariesExtHandler() *Dictionaries {
	return &Dictionaries{
		db: db.Init(),
	}
}

// 获取字典列表
func (d *Dictionaries) GetDictionaryTypes(limit, Offset int32) ([]*entity.DictionaryType, error) {
	d.db.SetTableName("dictionary_type");
	dtc := d.db.NewCondition()
	// 查询排除软删除
	dtc.SetFilter("delete", "0")
	// 分页
	pager := d.db.NewPager()
	// 设置每页显示条数
	pager.Limit = int(limit)
	// 设置偏移量
	pager.Offset = int(Offset)
	// 实现分页查询
	values, err := d.db.PagerFindAll(pager)
	if err != nil {
		return nil, err
	}
	var dictionaryTypes []*entity.DictionaryType
	values.Scan(dictionaryTypes)
	return dictionaryTypes, nil
}

// 添加字典类型
func (d *Dictionaries) AddDictionaryType(dictionaryType entity.AddDictionaryType) (int32, error) {
	d.db.SetTableName("dictionary_type");
	id, e := d.db.Insert(dictionaryType)
	if e != nil {
		return 0, e
	}
	return int32(id), nil
}

// 更新字典类型
func (d *Dictionaries) UpdateDictionaryType(dictionaryType entity.DictionaryType) (int32, error) {
	d.db.SetTableName("dictionary_type");
	id, e := d.db.Update(dictionaryType)
	if e != nil {
		return 0, e
	}
	return int32(id), nil
}

// 删除字典类型
func (d *Dictionaries) DeleteDictionaryType(id int32) (int32, error) {
	d.db.SetTableName("dictionary_type");
	// 软删除
	_, e := d.db.Update(entity.DelDictionaryType{
		Id:     id,
		Delete: true,
	})
	if e != nil {
		return 0, e
	}
	return int32(id), nil
}

// 添加字典
func (d *Dictionaries) AddDictionary(dTypeId int32, dictionary entity.AddDictionary) (int32, error) {
	return 0, nil
}

func (d *Dictionaries) UpdateDictionary() {

}
