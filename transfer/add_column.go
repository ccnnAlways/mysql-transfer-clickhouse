package transfer

import (
	"fmt"
	"mysql-transfer-clickhouse/database/clickhouse"
)

// getAllTables 获取所有的表名
func getAllTables() (tables []string, err error) {
	err = clickhouse.ClickhouseDB.Raw("show tables").Scan(&tables).Error
	return
}

// addColumn 给指定表添加指定字段
func addColumn(table, column, defaultV, commentV string) (err error) {
	sql := fmt.Sprintf("ALTER Table %s ADD column if not exists %s varchar(255)  default '%s' comment '%s';", table, column, defaultV, commentV)
	err = clickhouse.ClickhouseDB.Exec(sql).Error
	return
}

func addColumns(tables []string) (err error) {
	for _, t := range tables {
		err = addColumn(t, "base", "福建", "基地")
		if err != nil {
			break
		}
	}
	return
}

func ToAddColumns() (err error) {
	var tabels []string
	tabels, err = getAllTables()

	if err != nil {
		fmt.Println("添加字段发生异常 = ", err)
		return
	}

	if err = addColumns(tabels); err != nil {
		fmt.Println("添加字段发生异常 = ", err)
		return
	}
	return
}
