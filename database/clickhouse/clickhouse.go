package clickhouse

import (
	"fmt"

	"mysql-transfer-clickhouse/config"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var (
	ClickhouseDB *gorm.DB
)

func init() {
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s?dial_timeout=10s&read_timeout=20s", config.Conf.Clickhouse.User, config.Conf.Clickhouse.Passwd, config.Conf.Clickhouse.Host, config.Conf.Clickhouse.Port, config.Conf.Clickhouse.Database)
	var err error
	ClickhouseDB, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("打开Clickhouse数据库发生异常 = %v \n", err)
		return
	}
	fmt.Println("连接clickhouse 成功 ...")
}
