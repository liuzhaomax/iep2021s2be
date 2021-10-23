/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/3 3:17
 * @version     v1.0
 * @filename    db.go
 * @description
 ***************************************************************************/
package config

import (
	"cms/src/core"
	"cms/src/module.base/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logger "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func (cfg *Config) NewDB() (*gorm.DB, func(), error) {
	cfg.DB.DSN = cfg.Mysql.DSN()
	db, err := gorm.Open(cfg.DB.Type, cfg.DB.DSN)
	if err != nil {
		return nil, nil, err
	}
	if cfg.DB.Debug {
		db = db.Debug()
	}
	clean := func() {
		err := db.Close()
		if err != nil {
			logger.Errorf("DB closing failed: %s", err.Error())
		}
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, clean, err
	}
	db.SingularTable(true)
	db.SetLogger(&GormLogger{})
	db.DB().SetMaxIdleConns(cfg.DB.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(cfg.DB.MaxLifetime) * time.Second)
	return db, clean, err
}

func (cfg *Config) AutoMigrate(db *gorm.DB) error {
	dbType := strings.ToLower(cfg.DB.Type)
	if dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	db = db.AutoMigrate(new(entity.User)) // lztodo
	cfg.createAdmin(db)
	core.GetInstanceOfContext().LastUserId = cfg.GetLastUserId(db)
	return db.Error
}

func (cfg *Config) createAdmin(db *gorm.DB) {
	var admin = entity.User{}
	db.First(&admin)
	if admin.UserEmail == "" {
		admin.UserEmail = cfg.Admin.UserEmail
		admin.UserNickName = cfg.Admin.UserNickName
		if cfg.Admin.UserPassword == "" {
			admin.UserPassword = cfg.Admin.UserEmail
		}
		admin.UserCreateTime = time.Now()
		db.NewRecord(admin)
		db.Create(&admin)
		db.NewRecord(admin)
	}
}

func (cfg *Config) GetLastUserId(db *gorm.DB) uint {
	var user = entity.User{}
	db.Last(&user)
	return user.UserId
}
