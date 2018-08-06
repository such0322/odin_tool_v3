package models

import (
	"fmt"

	"odin_tool_v3/libs/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db    *gorm.DB
	dbCfg struct {
		Type, Host, Name, User, Passwd, Path, SSLMode string
	}
)

func NewEngine() (err error) {
	getConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", dbCfg.User, dbCfg.Passwd, dbCfg.Host, dbCfg.Name)
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	return nil
}

func getConfig() {
	cfg := setting.Cfg.Section("database")
	dbCfg.Type = cfg.Key("DB_TYPE").String()
	dbCfg.Host = cfg.Key("DB_HOST").String()
	dbCfg.Name = cfg.Key("DB_NAME").String()
	dbCfg.User = cfg.Key("DB_USER").String()
	dbCfg.Passwd = cfg.Key("DB_PASSWD").String()

}
