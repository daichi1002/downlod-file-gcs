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

	if err != nil {
		s.logger.Fatalf("file import error %v", err)
	}

	if len(objects) == 0 {
		s.logger.Info("no acquired file")
		os.Exit(0)
	}

	result := make([]byte, 0)
	for _, obj := range objects {
		content, err := client.DownloadFile(ctx, gcsBucketName, obj.Name)
		if err != nil {
			s.logger.Fatalf("file download error %v", err)
		}
		result = append(result, content...)
	}
}
