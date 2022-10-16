package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var initMySQLDB sync.Once

// 为了不被程序在运行时恶意修改, 设置成私有变量
var config *Config

// 全局MySQL 客户端实例
var db *sql.DB

func C() *Config {
	return config
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "calf",
		Host: "127.0.0.1",
		Port: "8089",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func (app *App) HttpAdress() string {
	return fmt.Sprintf("%s:%s", app.Host, app.Port)
}

func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Password:    "123456",
		Database:    "calf-go",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的MySQL连接池, 需要池做一些规划配置
	// 控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制MySQL复用, 比如5, 最多运行5个来复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 这个和MySQL Server配置有关系, 必须小于Server配置
	// 一个连接用12h 换一个conn, 保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle 连接 最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`

	lock sync.Mutex
}

// 1. 第一种方式, 使用LoadGlobal 在加载时 初始化全局db实例
// 2. 第二种方式, 惰性加载, 获取DB是，动态判断再初始化
func (m *MySQL) GetDB() *sql.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}
	return db
}

func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

func NewDefaultLog() *Log {
	return &Log{
		// debug, info, error, warn
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

// Log todo
// 用于配置全局Logger对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
}
