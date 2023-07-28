package main

import "mysql-transfer-clickhouse/transfer"

func main() {
	transfer.GenTransferCode()

	// // 迁移表结构
	// transfer.Transfer()

	// // 每张表添加特定字段
	// transfer.ToAddColumns()
}
