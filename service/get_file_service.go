package service

import (
	"context"
	"downlod-file-gcs/util"
)

type Service struct{}

// NewGetFileService の初期化処理
func NewGetFileService() *Service {
	// サービスの生成
	service := new(Service)
	return service
}

var logger = util.NewLogger()

func (service *Service) Execute(ctx context.Context) {
	logger.Info("ファイルダウンロード処理開始")
	logger.Info("ファイルダウンロード処理終了")
}
