package service

import (
	"context"
	"downlod-file-gcs/interfaces"
	"downlod-file-gcs/util"

	"gorm.io/gorm"
)

type Service struct{}

// NewGetFileService の初期化処理
func NewGetFileService() interfaces.Service {
	// サービスの生成
	service := new(Service)
	return service
}

var logger = util.NewLogger()

func (service *Service) Execute(ctx context.Context, db *gorm.DB, gcsClient interfaces.GcsClient) {
	logger.Info("start file download")
	logger.Info("end file download")
}
