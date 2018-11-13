package handler

import (
	"auditIntegral/_public/config"
	"auditIntegral/_public/log"
	"auditIntegral/systemSetup/db"
	"context"
	"go.uber.org/zap"

	dictionaries "auditIntegral/systemSetup/proto/dictionaries"
)

type Dictionaries struct {
	logger *zap.Logger
	db     *db.Dictionaries
}

func NewSystemSetupServiceExtHandler() *Dictionaries {
	return &Dictionaries{
		logger: log.Instance(),
		db:     db.NewDictionariesExtHandler(),
	}
}

// 列表
func (d *Dictionaries) List(ctx context.Context, req *dictionaries.ListRequest, rsp *dictionaries.ListResponse) error {
	dictionaryTypes, err := d.db.GetDictionaryTypes(1, 10)
	if err != nil {
		d.logger.Error("")
	}
	d.logger.Info("res", zap.Any(config.SystemSetupNameSpace, dictionaryTypes))
	return nil
}

// 添加
func (d *Dictionaries) Add(ctx context.Context, req *dictionaries.AddDictionaryType, rsp *dictionaries.EditResponse) error {
	//log.Printf("req: %v", req)
	return nil
}

// 修改
func (d *Dictionaries) Edit(ctx context.Context, req *dictionaries.DictionaryType, rsp *dictionaries.EditResponse) error {
	//log.Printf("req: %v", req)
	return nil
}

// 删除
func (d *Dictionaries) Delete(ctx context.Context, req *dictionaries.DelRequest, rsp *dictionaries.EditResponse) error {
	//log.Printf("dictionaries Delete req: %v", req)
	//log.Printf("ids: %v", req.GetIds())
	rsp.Status = &dictionaries.Status{
		Code:    0,
		Success: true,
		Msg:     req.Ids,
	}
	return nil
}
