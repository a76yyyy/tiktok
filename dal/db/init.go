// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"fmt"
	"time"

	"github.com/a76yyyy/tiktok/pkg/dlog"
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	DB     *gorm.DB
	Config = ttviper.ConfigInit("TIKTOK_DB", "dbConfig")
)

// Init init DB
func InitDB() {
	var err error

	logger := zapgorm2.New(dlog.InitLog())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	viper := Config.Viper
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.charset"),
		viper.GetBool("mysql.parseTime"),
		viper.GetString("mysql.loc"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			// PrepareStmt: true,
			// SkipDefaultTransaction: true,
			Logger: logger,
		},
	)
	if err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	if err = DB.Use(otelgorm.NewPlugin()); err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	if err := DB.AutoMigrate(&User{}, &Video{}, &Comment{}, &Relation{}); err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		logger.ZapLogger.Panic(err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Init() {
	InitDB()
}
