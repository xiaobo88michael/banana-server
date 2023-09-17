package config

// import (
// 	"fmt"
// 	"sync"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var DB DataBases

// type DataBases struct {
// 	MysqlDB mysqlDB
// }

// type mysqlDB struct {
// 	sync.RWMutex
// 	db *gorm.DB
// }

// type Writer struct{}

// func (w Writer) Printf(format string, args ...interface{}) {
// 	fmt.Printf(format, args...)
// }

// func InitMysqlDB() {
// 	fmt.Println("init mysqlDB start")
// 	//When there is no open IM database, connect to the mysql built-in database to create openIM database
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
// 		BoxDataConfig.Mysql.DBUserName, BoxDataConfig.Mysql.DBPassword, BoxDataConfig.Mysql.DBAddress[0], "mysql")
// 	var db *gorm.DB
// 	var err1 error
// 	db, err := gorm.Open(mysql.Open(dsn), nil)
// 	if err != nil {
// 		fmt.Println("Open failed ", err.Error(), dsn)
// 	}
// 	if err != nil {
// 		time.Sleep(time.Duration(30) * time.Second)
// 		db, err1 = gorm.Open(mysql.Open(dsn), nil)
// 		if err1 != nil {
// 			fmt.Println("Open failed ", err1.Error(), dsn)
// 			panic(err1.Error())
// 		}
// 	}

// 	//Check the database and table during initialization
// 	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;", BoxDataConfig.Mysql.DBDatabaseName)
// 	fmt.Println("exec sql: ", sql, " begin")
// 	err = db.Exec(sql).Error
// 	if err != nil {
// 		fmt.Println("Exec failed ", err.Error(), sql)
// 		panic(err.Error())
// 	}
// 	fmt.Println("exec sql: ", sql, " end")
// 	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
// 		BoxDataConfig.Mysql.DBUserName, BoxDataConfig.Mysql.DBPassword, BoxDataConfig.Mysql.DBAddress[0], BoxDataConfig.Mysql.DBDatabaseName)

// 	newLogger := logger.New(
// 		Writer{},
// 		logger.Config{
// 			SlowThreshold:             time.Duration(BoxDataConfig.Mysql.SlowThreshold) * time.Millisecond, // Slow SQL threshold
// 			LogLevel:                  logger.LogLevel(BoxDataConfig.Mysql.LogLevel),                       // Log level
// 			IgnoreRecordNotFoundError: true,                                                                // Ignore ErrRecordNotFound error for logger
// 			Colorful:                  true,                                                                // Disable color
// 		},
// 	)
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: newLogger,
// 	})
// 	if err != nil {
// 		fmt.Println("Open failed ", err.Error(), dsn)
// 		panic(err.Error())
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(BoxDataConfig.Mysql.DBMaxLifeTime))
// 	sqlDB.SetMaxOpenConns(BoxDataConfig.Mysql.DBMaxOpenConns)
// 	sqlDB.SetMaxIdleConns(BoxDataConfig.Mysql.DBMaxIdleConns)

// 	fmt.Println("open mysql ok ", dsn)

// 	DB.MysqlDB.db = db
// 	return
// }

// func (m *mysqlDB) DefaultGormDB() *gorm.DB {
// 	return DB.MysqlDB.db
// }
