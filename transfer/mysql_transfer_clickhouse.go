package transfer

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"mysql-transfer-clickhouse/database/mysql"
	"os"
	"os/exec"
)

var (
	modelPath = "model"
)

// 生成迁移的代码
func GenTransferCode() {
	// 生成model文件
	if err := exeCMD(); err != nil {
		fmt.Println("执行cmd 发生异常 = ", err)
		return
	}
	// 得到 model下的go文件 切片
	files, err := getFiles()

	if err != nil {
		fmt.Println("获取model文件列表发生异常 = ", err)
		return
	}
	// 处理得到的go文件切片，得到其中的所有结构体切片
	structs, err := getStructs(files)

	if err != nil {
		fmt.Println("获取model结构体列表发生异常 = ", err)
		return
	}

	// 处理结构体切片，生成代码
	writeFile(structs)
}

// 将内容写入文件
func writeFile(structs []string) error {
	var fileName = "transfer/transfer.go"

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	w.WriteString(`package transfer

import (
	"mysql-transfer-clickhouse/database/clickhouse"
	"mysql-transfer-clickhouse/model"
)
	
func Transfer() {`)

	for _, s := range structs {
		w.WriteString(fmt.Sprintf("\n\t clickhouse.ClickhouseDB.AutoMigrate(&model.%s{})", s))
	}

	w.WriteString("\n }")
	w.Flush()
	return nil
}

// 获取文件夹下的所有文件
func getFiles() (files []fs.FileInfo, err error) {
	files, err = ioutil.ReadDir(modelPath)
	return
}

// 获取所有的文件中的所有结构体
func getStructs(files []fs.FileInfo) (structs []string, err error) {

	for _, f := range files {
		var tmpStructs []string = make([]string, 0)
		tmpStructs, err = getStructFromFile(f.Name())

		if err != nil {
			return
		}
		structs = append(structs, tmpStructs...)
	}
	return
}

// getStructFromFile 从文件中获取所有的结构体名称
func getStructFromFile(fileName string) (structs []string, err error) {

	f, _ := ioutil.ReadFile(modelPath + "/" + fileName)
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, fileName, f, parser.ParseComments)
	if err != nil {
		return
	}

	for _, v := range p.Decls {
		if stc, ok := v.(*ast.GenDecl); ok && stc.Tok == token.TYPE {

			for _, spec := range stc.Specs {
				if tp, ok := spec.(*ast.TypeSpec); ok {
					structName := tp.Name.Name
					structs = append(structs, structName)
				}
			}
		}
	}
	return
}

// 执行命令行，在model文件夹下生成model文件
func exeCMD() error {
	dsn := mysql.GetMsqlDsn()

	cmd := exec.Command("gentool", "-dsn", dsn, "-onlyModel", "-outPath", modelPath)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("执行生成生成model struct 发生异常 = ", err)
		return err
	}

	logPath := "log/cmd.log"

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("创建日志发生异常 = ", err)
		return err
	}

	defer f.Close()

	if err = ioutil.WriteFile(logPath, output, 0666); err != nil {
		fmt.Println("cmd 日志写入文件发生异常 = ", err)
		return err
	}

	return err
}
