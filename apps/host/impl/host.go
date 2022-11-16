package impl

import (
	"calf-restful-api/apps"
	"calf-restful-api/apps/host"
	"calf-restful-api/conf"
	"database/sql"
	"github.com/chenjiahao-stl/go-framework/logger"
	"github.com/chenjiahao-stl/go-framework/logger/zap"
)

var impl = &IHostServiceImpl{}

type IHostDao struct {
	db *sql.DB
	l  logger.Logger
}

type IHostServiceImpl struct {
	dao *IHostDao
	l   logger.Logger
}

func (i *IHostServiceImpl) Config() {
	//注入db
	db := conf.C().MySQL.GetDB()
	dao := &IHostDao{
		db: db,
		l:  zap.L().Named("Host Dao"),
	}
	i.l = zap.L().Named("Host Service")
	i.dao = dao
}

func (i *IHostServiceImpl) Name() string {
	return host.APPNAME
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.RegistryImpl(impl)
}
