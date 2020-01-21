package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	common "github.com/bilibili/kratos/tool/pkg"
)

func main() {
	log.SetFlags(0)
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 64*1024)
			buf = buf[:runtime.Stack(buf, false)]
			log.Fatalf("程序解析失败, err: %+v stack: %s", err, buf)
		}
	}()

	code := common.FormatCode("test\ndemo")
	// Write to file.
	dir := filepath.Dir(".")
	outputName := filepath.Join(dir, "test.model.go")
	err := ioutil.WriteFile(outputName, []byte(code), 0644)
	if err != nil {
		log.Fatalf("写入文件失败: %s", err)
	}
	log.Println("test.model.go: 生成成功")
}
