package service

import "context"

type Service struct{}

// NewGetFileService の初期化処理
func NewGetFileService() *Service {
	// サービスの生成
	service := new(Service)
	return service
}

func (service *Service) Execute(ctx context.Context) {}
