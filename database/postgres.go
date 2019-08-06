package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *gorm.DB

// InitDatabase 初始化数据库
func InitDatabase() {

	if viper.GetString("database.type") != "postgres" {
		logrus.Fatal("不支持的数据库类型")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=3 client_encoding=UTF8",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
	)

	db, err := gorm.Open(viper.GetString("database.type"), dsn)
	if err != nil {
		logrus.Fatalf("数据库连接失败: %v", err)
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.max_open_conns"))

	if viper.GetString("mode") == "debug" {
		db.LogMode(true)
	}

	if err = db.DB().Ping(); err != nil {
		logrus.Fatalf("数据库心跳检测失败: %v", err)
	}
	DB = db

}

func Close() error {
	return DB.Close()
}
