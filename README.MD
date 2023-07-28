### Mysql 转移表结构到Clickhouse



#### 总体思路

* 使用gentool，逆向读取数据库，生成model的struct
* 然后再通过Gorm连接clickhouse，生成表结构





#### 测试

结果：通过。

环境：

* Windows

* Go：go1.19.4 windows/amd64

测试方法：

* 通过gentool工具对mysql的表结构生成model文件（`gentool -dsn "root:xxxxpwd@tcp(127.0.0.1:3306)/xxx?charset=utf8mb4&parseTime=True&loc=Local" -tables "tb_accomplish,tb_annunciate" -onlyModel -outPath "./model"`）
* 使用gorm连接clickhouse，生成表。（`clickhouse.ClickhouseDB.AutoMigrate(&model.TbAccomplish{})`）





#### 思路

* 通过yaml配置
  * mysql（用户、密码、数据库、地址，端口）
  * clickhouse（用户名、密码、端口、数据库、地址）
* 通过gentool在临时目录：`./model`中生成model文件，
* 通过`ast`技术，得到所有的结构体的名称
* 生成在Clickhouse创建表结构的代码
* 执行



#### 执行

1. `go mod tidy`
2. 注释掉`main.go`中的其他方法，只保留`transfer.GenTransferCode()`

3. 配置配置文件`etc/config.yml`
4. `go run main.go`
5. 注释掉 `transfer.GenTransferCode()`，解开迁移的代码注释`transfer.Transfer()`
6. `go run main.go`



#### 后续

- [ ] 直接执行`go run main.go`即可迁移数据结构
