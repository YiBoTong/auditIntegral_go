package handler

import (
	dictionaries "auditIntegral/systemSetup/proto/dictionaries"
	"context"
	"log"
)

type Dictionaries struct{}

// 列表
func (d *Dictionaries) List(ctx context.Context, req *dictionaries.ListRequest, rsp *dictionaries.ListResponse) error {
	log.Fatalf("req: %v", req)
	return nil
}

// 添加
func (d *Dictionaries) Add(ctx context.Context, req *dictionaries.AddDictionaryType, rsp *dictionaries.EditResponse) error {
	log.Printf("req: %v", req)
	return nil
}

// 修改
func (d *Dictionaries) Edit(ctx context.Context, req *dictionaries.DictionaryType, rsp *dictionaries.EditResponse) error {
	log.Printf("req: %v", req)
	return nil
}

// 删除
func (d *Dictionaries) Delete(ctx context.Context, req *dictionaries.DelRequest, rsp *dictionaries.EditResponse) error {
	log.Printf("dictionaries Delete req: %v", req)
	log.Printf("ids: %v", req.GetIds())
	var Status dictionaries.Status
	rsp.Data = "成功"
	Status.Success = true
	Status.Code = 1
	Status.Msg = "none"
	rsp.Status = &Status
	return nil
}
