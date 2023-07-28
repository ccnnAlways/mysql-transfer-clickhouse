package mysql

import (
	"fmt"

	"mysql-transfer-clickhouse/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn     string
	MysqlDB *gorm.DB
)

func init() {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Mysql.User, config.Conf.Mysql.Passwd, config.Conf.Mysql.Host, config.Conf.Mysql.Port, config.Conf.Mysql.Database)

	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("打开Mysql发生异常 = %v \n", err)
		return
	}
	fmt.Println("Mysql 成功 ...")
}

// GetMsqlDsn 获取Mysql DSN
func GetMsqlDsn() string {
	return dsn
}
