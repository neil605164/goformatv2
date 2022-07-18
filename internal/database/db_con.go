package database

import (
	"fmt"
	"goformatv2/app/global"
	"goformatv2/app/global/errorcode"
	"goformatv2/app/global/helper"
	"goformatv2/app/model"
	"log"
	"os"
	"time"

	"gorm.io/plugin/dbresolver"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// dbCon DB連線資料
type dbCon struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var DB *gorm.DB

func DBConn() (*gorm.DB, errorcode.Error) {

	var err error

	dsnMaster := composeString(global.DBMaster)
	dsnSlave := composeString(global.DBSlaver)

	// 連接gorm
	DB, err = gorm.Open(mysql.Open(dsnMaster), &gorm.Config{})
	if err != nil {
		apiErr := helper.ErrorHandle(global.FatalLog, errorcode.Code.DBConnectError, err.Error())
		return nil, apiErr
	}

	_ = DB.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsnMaster)},
			Replicas: []gorm.Dialector{mysql.Open(dsnSlave)},
			Policy:   dbresolver.RandomPolicy{}}).
			// 空閒連線 timeout 時間
			SetConnMaxIdleTime(15 * time.Second).
			// 空閒連線 timeout 時間
			SetConnMaxLifetime(15 * time.Second).
			// 限制最大閒置連線數
			SetMaxIdleConns(100).
			// 限制最大開啟的連線數
			SetMaxOpenConns(2000),
	)

	sqlDB, _ := DB.DB()

	// 限制最大閒置連線數
	sqlDB.SetMaxIdleConns(100)
	// 限制最大開啟的連線數
	sqlDB.SetMaxOpenConns(2000)
	// 空閒連線 timeout 時間
	sqlDB.SetConnMaxLifetime(15 * time.Second)

	if global.Config.DB.Debug {
		DB = DB.Debug()
	}

	return DB, nil
}

// DBPing 檢查DB是否啟動
func DBPing() {
	// 檢查 master db
	dbCon, apiErr := DBConn()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBMaster.Host)
	}

	masterDB, err := dbCon.DB()
	if err != nil {
		log.Fatalf("🔔🔔🔔 CONNECT MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}

	err = masterDB.Ping()
	if err != nil {
		log.Fatalf("🔔🔔🔔 PING MASTER DB ERROR: %v 🔔🔔🔔", err.Error())
	}

}

// CheckTable 啟動main.go服務時，直接檢查所有 DB 的 Table 是否已經存在
func CheckTable() {
	db, apiErr := DBConn()
	if apiErr != nil {
		log.Fatalf("🔔🔔🔔 MASTER DB CONNECT ERROR: %v 🔔🔔🔔", global.Config.DBMaster.Host)
	}

	// 會自動建置 DB Table
	err := db.Set("gorm:table_options", "comment '細單規則'").AutoMigrate(&model.Admin{})
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		_ = helper.ErrorHandle(global.FatalLog, errorcode.Code.DBTableNotExist, fmt.Sprintf("❌ 設置DB錯誤： %v ❌", err.Error()))
		log.Fatalf("🔔🔔🔔 MIGRATE MASTER TABLE ERROR: %v 🔔🔔🔔", err.Error())
	}

}

// composeString 組合DB連線前的字串資料
func composeString(mode string) string {
	db := dbCon{}

	switch mode {
	case global.DBMaster:
		db.Host = getHost(true)
		db.Username = getUser(true)
		db.Password = getPass(true)
		db.Database = getDBName()
	case global.DBSlaver:
		db.Host = getHost(false)
		db.Username = getUser(false)
		db.Password = getPass(false)
		db.Database = getDBName()
	}

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?timeout=5s&charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Database)
}

// getHost 取 DB Host
func getHost(master bool) string {

	switch master {
	case true:
		if value, ok := os.LookupEnv("MHOST"); ok {
			return value
		}
		return global.Config.DBMaster.Host
	default:
		if value, ok := os.LookupEnv("SHOST"); ok {
			return value
		}
		return global.Config.DBMaster.Host
	}
}

// getUser 取 DB User
func getUser(master bool) string {

	switch master {
	case true:
		if value, ok := os.LookupEnv("MUSER"); ok {
			return value
		}
		return global.Config.DBMaster.Username
	default:
		if value, ok := os.LookupEnv("SUSER"); ok {
			return value
		}
		return global.Config.DBMaster.Username
	}
}

// getPass 取 DB Pass
func getPass(master bool) string {

	switch master {
	case true:
		if value, ok := os.LookupEnv("MPASS"); ok {
			return value
		}
		return global.Config.DBMaster.Password
	default:
		if value, ok := os.LookupEnv("SPASS"); ok {
			return value
		}
		return global.Config.DBMaster.Password
	}
}

// getDBName 取 DB Name
func getDBName() string {

	if value, ok := os.LookupEnv("DB_NAME"); ok {
		return value
	}
	return global.Config.DBMaster.Database
}
