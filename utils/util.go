package utils

import (
	"fmt"
	"github.com/teris-io/shortid"
	"reflect"
	"runtime"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetLineInfo() () {
	//pc 计数器， file 文件名， line 行号， ok 是否
	// runtime.Caller(4)这里的4是一个层级关系，可以尝试使用0 1 2 3来看看
	// 4 在其他项目中使用的时候，如果在log的test中，使用3
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fmt.Println(file)
		fmt.Println(runtime.FuncForPC(pc).Name())
		fmt.Println(line)
	}
}

func NameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
