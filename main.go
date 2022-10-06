package main

import (
	"downlod-file-gcs/util"
	"flag"
	"fmt"
)

var logger = util.NewLogger()

func main() {
	// 起動引数の処理
	reqProcessingDate := parseArgs()
	fmt.Println(reqProcessingDate)
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
