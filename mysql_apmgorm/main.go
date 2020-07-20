package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	gormzap "github.com/wantedly/gorm-zap"
	"go.elastic.co/apm/module/apmgorm"
	_ "go.elastic.co/apm/module/apmgorm/dialects/mysql"
	"go.uber.org/zap"
)

// NewClient factory methord of *gorm.DB
func NewClient(user, pwd, host, dbname string, logJSON bool) (*gorm.DB, error) {
	db, err := apmgorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", user, pwd, host, dbname))
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	if logJSON {
		db.SetLogger(gormzap.New(zap.L()))
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
	return db, nil
}
