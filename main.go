package main

import (
	"context"
	"downlod-file-gcs/config"
	"downlod-file-gcs/constant"
	"downlod-file-gcs/service"
	"downlod-file-gcs/util"
	"flag"

	"gorm.io/gorm"
)

var logger = util.NewLogger()

func main() {
	// 起動引数の処理
	reqProcessingDate := parseArgs()

	logger.Infof("%sのバッチ処理開始", reqProcessingDate)

	// contextに値を設定
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.ProcessingDateContextKey, reqProcessingDate)

	// DB接続処理
	databaseConfig := config.GetDatabaseConfig()
	db, err := databaseConfig.ConnectDatabaseWithGorm(10)
	if err != nil {
		logger.Fatalf("Failed to connect database")
	}
	defer closeGormDB(db)

	// service初期化処理
	service := service.NewGetFileService()

	// バッチ処理実行
	service.Execute(ctx)
}

func parseArgs() string {
	// 処理日時を取得
	reqProcessingDate := flag.String("date", "", "処理日時")

	flag.Parse()

	// 処理日時をフォーマットする
	processingDate, err := util.CreateProcessingDate(*reqProcessingDate)

	if err != nil {
		logger.Fatalf("日付のparseに失敗しました", err)
	}

	return processingDate
}

func closeGormDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("failed to extract sql db from gorm.DB instance, error: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		logger.Errorf("failed to close sql db, error: %v", err)
	}
}
