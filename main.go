package main

import (
	"embed"
	"goformatv2/app/global"
	"goformatv2/app/global/helper"
	"goformatv2/internal/cache"
	"goformatv2/internal/database"
	"goformatv2/internal/entry"
	"os"

	_ "goformatv2/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

//go:embed env/*
var f embed.FS

// 初始化動作
func init() {

	// 載入環境設定，所有動作須在該func後執行
	global.Start(f)

	// 設定 log 格式
	helper.SetFormatter(&logrus.JSONFormatter{})

	// 檢查 DB 機器服務
	db := database.NewDBConnection()
	db.DBPing()

	// 自動建置 DB + Table
	if helper.IsLocal() {
		db.CheckTable()
	}

	// 檢查 Redis 機器服務
	redis := cache.RedisIns()
	_ = redis.Ping()

	// 設定程式碼 timezone
	os.Setenv("TZ", "America/Puerto_Rico")
}

func main() {
	entry.Run()
}
