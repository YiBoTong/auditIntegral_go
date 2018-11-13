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
	dictionaryTypeCondition := d.db.NewCondition()
	// 查询排除软删除
	dictionaryTypeCondition.SetFilter("delete", "0")
	// 分页
	pager := d.db.NewPager()
	// 设置每页显示条数
	pager.Limit = limit
	// 设置偏移量
	pager.Offset = Offset
	// 实现分页查询
	values, err := d.db.PagerFindAll(pager)
	if err != nil {
		return nil, err
	}
	var dictionaryTypes []*entity.DictionaryType
	values.Scan(dictionaryTypes)
	return dictionaryTypes, nil
}
