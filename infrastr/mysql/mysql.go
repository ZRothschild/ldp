package mysql

import (
	"github.com/ZRothschild/ldp/app/user/userM"
	"github.com/ZRothschild/ldp/infrastr/log"
	"github.com/ZRothschild/ldp/infrastr/static/config"
	mysqlD "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type (
	DB struct {
		*gorm.DB
	}
)

func NewDb(cfg *config.Config) *DB {
	var (
		db = new(DB)
	)
	if err := db.Open(cfg); err != nil {
		//panic(err)
	}

	return db
}

func (db *DB) NewConfig(cfg *config.Config) *mysqlD.Config {
	var (
		c = mysqlD.NewConfig()
	)
	c.User = cfg.Mysql.User
	c.Passwd = cfg.Mysql.Passwd
	c.Net = cfg.Mysql.Net
	c.Addr = cfg.Mysql.Addr
	c.DBName = cfg.Mysql.DBName
	c.Collation = cfg.Mysql.Collation
	c.Loc = time.Local
	return c
}

func (db *DB) Open(cfg *config.Config) error {
	var (
		err error
	)
	if db.DB, err = gorm.Open(mysql.New(mysql.Config{
		//DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DSNConfig:                 db.NewConfig(cfg),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		DisableAutomaticPing: true,
		Logger: log.NewDefault(log.SLog, logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		}),
	}); err != nil {
		return err
	}

	if pinger, ok := db.ConnPool.(interface{ Ping() error }); ok {
		err = pinger.Ping()
	}

	if err = db.Migrator(); err != nil {
		return err
	}
	var (
		u = make([]userM.User, 0)
	)
	db.Model(userM.User{}).Find(&u)
	return err
}

func (db *DB) Migrator() error {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';").AutoMigrate(&userM.User{}); err != nil {
		return err
	}
	return nil
}
