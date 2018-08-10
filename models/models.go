package models

import (
	"fmt"
	"log"
	"odin_tool_v3/libs/setting"
	AM "odin_tool_v3/models/api/master"
	BMi "odin_tool_v3/models/bridge/misc"
	"odin_tool_v3/models/tool"

	"odin_tool_v3/models/api/master/bonus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	manager = &DBManager{}
	dbCfgs  = map[string]*dbCfg{}
)

type DBManager struct {
	DBS map[string]*gorm.DB
}

func (m *DBManager) getDB(dbName string) *gorm.DB {
	fmt.Println("DBS:", m.DBS[dbName])
	db, ok := m.DBS[dbName]
	if !ok {
		var section string
		switch dbName {
		case "tool":
			section = "database"
		case "api_master":
			section = "database_api_master"

		}
		getConfig(section, dbName)
		var err error
		db, err = m.newEngine(dbName)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return db
}

func (m *DBManager) newEngine(dbName string) (db *gorm.DB, err error) {
	dbCfg := dbCfgs[dbName]
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", dbCfg.User, dbCfg.Passwd, dbCfg.Host, dbCfg.Name)
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	m.DBS[dbName] = db
	return db, nil
}

func UseDB(dbName string) *gorm.DB {
	fmt.Println("use DB:", dbName)
	db := manager.getDB(dbName)
	return db
}

type dbCfg struct {
	Type, Host, Name, User, Passwd, Path, SSLMode string
}

func NewEngines() {
	//new 一个数据库manager
	manager.DBS = make(map[string]*gorm.DB)

	getConfig("database", "tool")
	tool.DB, _ = manager.newEngine("tool")

	getConfig("database_api_master", "api_master")
	AM.DB, _ = manager.newEngine("api_master")
	bonus.DB = AM.DB

	getConfig("database_bridge_misc", "bridge_misc")
	BMi.DB, _ = manager.newEngine("bridge_misc")
}

func getConfig(section, dbName string) {
	cfg := setting.Cfg.Section(section)
	dbCfg := &dbCfg{}
	dbCfg.Type = cfg.Key("DB_TYPE").String()
	dbCfg.Host = cfg.Key("DB_HOST").String()
	dbCfg.Name = cfg.Key("DB_NAME").String()
	dbCfg.User = cfg.Key("DB_USER").String()
	dbCfg.Passwd = cfg.Key("DB_PASSWD").String()
	dbCfgs[dbName] = dbCfg
}
