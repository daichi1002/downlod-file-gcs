package service

import (
	"context"
	"downlod-file-gcs/interfaces"
	"downlod-file-gcs/util"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger *zap.SugaredLogger
}

// NewGetFileService の初期化処理
func NewGetFileService() interfaces.Service {

	logger := util.NewLogger()
	// サービスの生成
	service := &Service{
		logger: logger,
	}
	return service
}

func (s *Service) Execute(ctx context.Context, db *gorm.DB, gcsClient interfaces.GcsClient) {
	s.logger.Info("start file download")
	s.DownloadFile(ctx, gcsClient)
	s.logger.Info("end file download")
}

func (s *Service) DownloadFile(ctx context.Context, client interfaces.GcsClient) {
	// GCSバケット指定
	gcsBucketName := os.Getenv("GcsBacketName")
	date := time.Now().Format("20060102")
	prefix := fmt.Sprintf("%v/", date)

	objects, err := client.ListFilesWithPrefix(ctx, gcsBucketName, prefix)
	fmt.Println(objects)
	if err != nil {
		s.logger.Fatalf("ファイル取り込みエラーが発生しました。%v", err)
	}
}
