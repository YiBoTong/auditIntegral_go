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
func (d *Dictionaries) GetDictionaryTypes(limit, Offset int) (*gosql.Pager, []*entity.DictionaryType, error) {
	d.db.SetTableName("dictionary_type")
	dtc := d.db.NewCondition()
	// 查询排除软删除
	dtc.SetFilter("delete", false)
	// 分页
	pager := d.db.NewPager()
	// 设置每页显示条数
	pager.Limit = int(limit)
	// 设置偏移量
	pager.Offset = int(Offset)
	// 实现分页查询
	values, err := d.db.SetCondition(dtc).PagerFindAll(pager)
	if err != nil {
		return nil, nil, err
	}
	var dictionaryTypes []*entity.DictionaryType
	values.Scan(&dictionaryTypes)
	return pager, dictionaryTypes, nil
}

// 添加字典类型
func (d *Dictionaries) AddDictionaryType(dictionaryType entity.AddDictionaryType) (int, error) {
	d.db.SetTableName("dictionary_type")
	id, e := d.db.Insert(dictionaryType)
	if e != nil {
		return 0, e
	}
	return int(id), nil
}

// 更新字典类型
func (d *Dictionaries) UpdateDictionaryType(dictionaryType entity.DictionaryType) (int, error) {
	d.db.SetTableName("dictionary_type")
	c := d.db.NewCondition().SetFilter("id", dictionaryType.Id)
	id, e := d.db.SetCondition(c).Update(dictionaryType)
	return id, e
}

// 删除字典类型
func (d *Dictionaries) DeleteDictionaryType(id int) (int, error) {
	d.db.SetTableName("dictionary_type")
	c := d.db.NewCondition().SetFilter("id", id)
	// 软删除
	_, e := d.db.SetCondition(c).Update(entity.DelDictionaryType{
		Id:     id,
		Delete: true,
	})
	return id, e
}

// 添加字典
func (d *Dictionaries) AddDictionary(dictionary entity.AddDictionary) (int, error) {
	d.db.SetTableName("dictionary")
	id, e := d.db.Insert(dictionary)
	return id, e
}

// 字典编辑
func (d *Dictionaries) UpdateDictionary(dictionary entity.Dictionary) (int, error) {
	d.db.SetTableName("dictionary")
	c := d.db.NewCondition().SetFilter("id", dictionary.Id)
	id, e := d.db.SetCondition(c).Update(dictionary)
	return id, e
}

// 字典删除
func (d *Dictionaries) DelDictionary(dictionaryTypeId int) (int, error) {
	d.db.SetTableName("dictionary")
	c := d.db.NewCondition().SetFilter("type_id", dictionaryTypeId)
	id, e := d.db.SetCondition(c).Del()
	return id, e
}
