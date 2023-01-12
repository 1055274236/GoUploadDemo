/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 19:44:43
 */
package mysql

import (
	"UploadApi/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func initDb() {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // 慢 SQL 阈值
	// 		LogLevel:                  logger.Info, // 日志级别
	// 		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	// 		Colorful:                  false,       // 禁用彩色打印
	// 	},
	// )

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=%v",
		config.DbUser, config.DbPassword, config.DbIp, config.DbPort, config.DbName, config.DbCharset, config.DbLoc)

	dbC, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger:                 newLogger,
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 缓存预编译语句
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true, // skip the snake_casing of names
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("数据库连接失败")
		log.Println(err)
		os.Exit(1)
	}
	log.Println("数据库连接成功")
	db = dbC

	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(config.DbMaxIdleConns) //设置最大连接数
	sqlDb.SetMaxOpenConns(config.DbMaxOpenConns) //设置最大的空闲连接数
}

func GetContentByTable(tableName string) *gorm.DB {
	if db == nil {
		initDb()
	}
	return db.Table(tableName)
}

func GetDB() *gorm.DB {
	if db == nil {
		initDb()
	}
	return db
}
