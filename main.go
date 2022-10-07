package main

import (
	"context"
	"downlod-file-gcs/constant"
	"downlod-file-gcs/util"
	"flag"
)

var logger = util.NewLogger()

func main() {
	// 起動引数の処理
	reqProcessingDate := parseArgs()

	logger.Infof("%sのバッチ処理開始", reqProcessingDate)

	// contextに値を設定
	ctx := context.Background()
	ctx = context.WithValue(ctx, constant.ProcessingDateContextKey, reqProcessingDate)
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
