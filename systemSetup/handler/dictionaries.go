package handler

import (
	"auditIntegral/_public/config"
	logs "auditIntegral/_public/log"
	"auditIntegral/_public/util"
	"auditIntegral/systemSetup/db"
	"auditIntegral/systemSetup/entity"
	dictionaries "auditIntegral/systemSetup/proto/dictionaries"
	"context"
	"errors"
	"go.uber.org/zap"
	"strconv"
)

type Dictionaries struct {
	logger *zap.Logger
	db     *db.Dictionaries
}

func NewDictionariesExtHandler() *Dictionaries {
	return &Dictionaries{
		logger: logs.Instance(),
		db:     db.NewDictionariesExtHandler(),
	}
}

// 列表
func (d *Dictionaries) List(ctx context.Context, req *dictionaries.ListRequest, rsp *dictionaries.ListResponse) error {
	limitNum := config.LimitNum // 每页显示数量
	offsetNum := 0              // 偏移量
	if req.Page.Size > 0 {
		limitNum = int(req.Page.Size)
	}
	if req.Page.Page > 0 {
		offsetNum = (int(req.Page.Page) - 1) * limitNum
	}

	dictionaryTypeArr := []*dictionaries.DictionaryType{}
	page, dictionaryTypes, err := d.db.GetDictionaryTypes(limitNum, offsetNum)
	if err != nil {
		d.logger.Error("[Dictionaries List]", zap.Error(err))
	} else {
		for _, v := range dictionaryTypes {
			dictionaryTypeArr = append(dictionaryTypeArr, &dictionaries.DictionaryType{
				Id:         int32(v.Id),
				TypeId:     int32(v.TypeId),
				Key:        v.Key,
				Title:      v.Title,
				IsUse:      v.IsUse,
				Describe:   v.Describe,
				UserId:     int32(v.UserId),
				UserName:   "测试",
				UpdateTime: v.UpdateTime,
			})
		}
	}
	error := err != nil
	rsp.Page = &dictionaries.ResponsePage{
		Size:  int32(page.Limit),
		Page:  int32(page.Offset/page.Limit + 1),
		Total: int32(page.Count),
	}
	rsp.Data = dictionaryTypeArr
	rsp.Status = &dictionaries.Status{
		Code:  0,
		Error: error,
		Msg:   config.GetTodoResMsg(config.ListStr, error),
	}
	return nil
}

// 添加
func (d *Dictionaries) Add(ctx context.Context, req *dictionaries.AddDictionaryType, rsp *dictionaries.EditResponse) error {
	dtId, e := d.db.AddDictionaryType(entity.AddDictionaryType{
		TypeId:     int(req.TypeId),
		Key:        req.Key,
		Title:      req.Title,
		IsUse:      req.IsUse,
		UpdateTime: util.GetLocalNowTimeStr(),
		UserId:     0,
		Describe:   req.Describe,
	})
	if e != nil {
		d.logger.Error("[Dictionaries Add]", zap.Error(e))
	} else {
		for _, dictionary := range req.Dictionaries {
			_, e := d.db.AddDictionary(entity.AddDictionary{
				TypeId:   dtId,
				Key:      dictionary.Key,
				Value:    dictionary.Value,
				Order:    int(dictionary.Order),
				Describe: dictionary.Describe,
			})
			if e != nil {
				break
			}
		}
	}
	error := e != nil
	rsp.Data = strconv.Itoa(dtId)
	rsp.Status = &dictionaries.Status{
		Code:  0,
		Error: error,
		Msg:   config.GetTodoResMsg(config.AddStr, error),
	}
	return nil
}

// 根据id或者key获取字典
func (d *Dictionaries) GetInfo(ctx context.Context, req *dictionaries.GetRequest, rsp *dictionaries.GetResponse) error {
	dtId := req.Id
	var dictionaryType *dictionaries.DictionaryType
	//var dictionaryArr []*dictionaries.Dictionary
	var err = errors.New("id is required")
	if dtId != 0 {
		var dictionaryTypeRes *entity.DictionaryType
		dictionaryTypeRes, err = d.db.GetDictionaryTypeById(int(dtId))
		dictionaryType = &dictionaries.DictionaryType{
			Id:         int32(dictionaryTypeRes.Id),
			TypeId:     int32(dictionaryTypeRes.TypeId),
			Key:        dictionaryTypeRes.Key,
			Title:      dictionaryTypeRes.Title,
			IsUse:      dictionaryTypeRes.IsUse,
			Describe:   dictionaryTypeRes.Describe,
			UserId:     int32(dictionaryTypeRes.UserId),
			UserName:   "--",
			UpdateTime: dictionaryTypeRes.UpdateTime,
		}
	}
	//if err == nil {
	//	var dictionaryArrRes []*entity.Dictionary
	//	dictionaryArrRes, err = d.db.GetDictionaryByDictionaryTypeId(int(dtId))
	//	for _, d := range dictionaryArrRes {
	//		dictionaryArr = append(dictionaryArr, &dictionaries.Dictionary{
	//			Id:       int32(d.Id),
	//			Key:      d.Key,
	//			Value:    d.Value,
	//			Order:    int32(d.Order),
	//			Describe: d.Describe,
	//		})
	//	}
	//	dictionaryType.Dictionaries = dictionaryArr
	//}
	if err != nil {
		d.logger.Error("[Dictionaries Get]", zap.Error(err))
	}
	error := err != nil
	rsp.Data = dictionaryType
	rsp.Status = &dictionaries.Status{
		Code:  0,
		Error: error,
		Msg:   config.GetTodoResMsg(config.GetStr, error),
	}
	return nil
}

// 修改
func (d *Dictionaries) Edit(ctx context.Context, req *dictionaries.DictionaryType, rsp *dictionaries.EditResponse) error {
	dtId := int(req.Id)
	_, e := d.db.UpdateDictionaryType(entity.DictionaryType{
		Id:         int(req.Id),
		TypeId:     int(req.TypeId),
		Key:        req.Key,
		Title:      req.Title,
		IsUse:      req.IsUse,
		UpdateTime: util.GetLocalNowTimeStr(),
		UserId:     0,
		Describe:   req.Describe,
	})
	// 硬删除
	_, err := d.db.DelDictionary(dtId)
	if e == nil && err == nil {
		for _, dictionary := range req.Dictionaries {
			// 重新添加
			_, e := d.db.AddDictionary(entity.AddDictionary{
				TypeId:   dtId,
				Key:      dictionary.Key,
				Value:    dictionary.Value,
				Order:    int(dictionary.Order),
				Describe: dictionary.Describe,
			})
			if e != nil {
				break
			}
		}
	} else {
		if e != nil {
			d.logger.Error("[Dictionaries Edit]", zap.Error(e))
		} else {
			d.logger.Error("[Dictionaries Edit]", zap.Error(err))
		}
	}
	error := e != nil && err != nil
	rsp.Data = strconv.Itoa(dtId)
	rsp.Status = &dictionaries.Status{
		Code:  0,
		Error: error,
		Msg:   config.GetTodoResMsg(config.EditStr, error),
	}
	return nil
}

// 删除
func (d *Dictionaries) Delete(ctx context.Context, req *dictionaries.DelRequest, rsp *dictionaries.EditResponse) error {
	error := false
	_, err := d.db.DeleteDictionaryType(int(req.Id))
	if err != nil {
		error = false
		d.logger.Error("[Dictionaries Del]", zap.Error(err))
	}
	rsp.Status = &dictionaries.Status{
		Code:  0,
		Error: error,
		Msg:   config.GetTodoResMsg(config.DelStr, error),
	}
	return nil
}
