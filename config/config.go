package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	config     Config
	dbInstance *gorm.DB
	cr         *cron.Cron
	Snowflake  *snowflake.Node
)

type Config struct {
	App   *App
	Mysql *Mysql
	Log   *Log
}

type App struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
	Port   string `yaml:"port"`
	Secret string `yaml:"secret"`
}

func (a *App) URL() string {
	return fmt.Sprintf("%v:%v", a.Server, a.Port)
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	lock     sync.Mutex
}

type Log struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
	To    string `yaml:"to"`
}

func (m *Mysql) Init(l *log.Logger) error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=UTC", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(l, logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("数据库连接错误：%v\n", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}

	dbInstance = db

	return nil
}

// Get 获取全局config
func Get() *Config {
	return &config
}

// GetDB 获取DB连接
func GetDB() *gorm.DB {
	return dbInstance
}

// GetCron 获取定时任务
func GetCron() *cron.Cron {
	return cr
}
