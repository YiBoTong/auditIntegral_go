package handler

import (
	"auditIntegral/_public/config"
	logs "auditIntegral/_public/log"
	"auditIntegral/_public/util"
	"auditIntegral/systemSetup/db"
	"auditIntegral/systemSetup/entity"
	dictionaries "auditIntegral/systemSetup/proto/dictionaries"
	"context"
	"go.uber.org/zap"
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
	dictionaryTypes, err := d.db.GetDictionaryTypes(req.Page.Size, req.Page.Page*req.Page.Size)
	if err != nil {
		d.logger.Error("")
	}
	d.logger.Info("res", zap.Any(config.SystemSetupNameSpace, dictionaryTypes))
	return nil
}

// 添加
func (d *Dictionaries) Add(ctx context.Context, req *dictionaries.AddDictionaryType, rsp *dictionaries.EditResponse) error {
	dtId, e := d.db.AddDictionaryType(entity.AddDictionaryType{
		TypeId:     req.TypeId,
		Key:        req.Key,
		Title:      req.Title,
		IsUse:      req.IsUse,
		UpdateTime: util.GetLocalNowTimeStr(),
		UserId:     0,
		Describe:   req.Describe,
	})
	if e != nil {
		d.logger.Error("[Dictionaries Add]", zap.Error(e))
		return e
	}
	for i, dictionary := range req.Dictionaries {
		d.logger.Info("add dictionary "+string(dtId), zap.Any("dictionaryItem"+string(i), dictionary))
	}
	return nil
}

// 修改
func (d *Dictionaries) Edit(ctx context.Context, req *dictionaries.DictionaryType, rsp *dictionaries.EditResponse) error {
	//log.Printf("req: %v", req)
	return nil
}

// 删除
func (d *Dictionaries) Delete(ctx context.Context, req *dictionaries.DelRequest, rsp *dictionaries.EditResponse) error {
	error := false
	_, err := d.db.DeleteDictionaryType(req.Id)
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
