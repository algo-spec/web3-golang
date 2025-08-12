package database

import (
	"task04/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect 连接数据库
func Connect() (*gorm.DB, error) {
	//加载配置
	config.LoadConfig()

	db, err := gorm.Open(mysql.Open(config.AppConfig.DBPath))
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	DB = db
	return DB, nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}

// Close 关闭数据库连接
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

