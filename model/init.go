package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// 以下为配置mysql数据库
var (
	addrMYSQL = "127.0.0.1:3306" // mysql地址
	account	  = "root" 			 // mysql账号
	password  = "xbgydx138386"	 // mysql密码
	dbName	  = "bytedancecamp"  // mysql数据库
)

var db *gorm.DB

// InitMysql 初始化mysql链接
func InitMysql() {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 				time.Second, // Slow SQL threshold 慢查询时间
			LogLevel: 					logger.Info, // Log level(这里要根据需求修改一下）
			IgnoreRecordNotFoundError: 	true, // Ignore ErrRecordNotFound error for logger
			Colorful: 					false, // Disable color
		},
	)
	connString := account + ":" + password + "@tcp(" + addrMYSQL + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dB, err := gorm.Open(mysql.Open(connString), &gorm.Config{ // gorm.Open:打开初始化数据库会话基于dialector
		Logger:					newLogger,
		SkipDefaultTransaction: false, /* 自动开启事务的开关 GORM默认在事务中执行单一的创建、更新、删除操作，以确保数据库数据的完整性。
		* 您可以通过设置“SkipDefaultTransaction”为true来禁用它 */
	})
	sqlDB, err := dB.DB()
	if err != nil {
		log.Fatalln("mysql lost:", err)
	}
	// 设置连接池
	// 空闲
	sqlDB.SetMaxIdleConns(10) /* 设置空闲连接池的最大连接数。
	如果MaxOpenConns大于0但小于新的MaxIdleConns，那么新的MaxIdleConns将被减少以匹配MaxOpenConns的限制。
	n <= 0表示不保留空闲连接 */
	// 打开
	sqlDB.SetMaxOpenConns(30) /* SetMaxOpenConns设置数据库的最大打开连接数。
	如果MaxIdleConns大于0并且新的MaxIdleConns小于MaxIdleConns，那么MaxIdleConns将会减少以匹配新的MaxOpenConns限制。
	如果n <= 0，则对打开连接的数量没有限制。默认值是0(无限) */
	db = dB
}
