/*
创建时间: 2019/11/6
作者: zjy
功能介绍:
工具包
*/

package xutil

import (
	"encoding/csv"
	"fmt"
	"io"
	"../xlog"
	"os"
)


func MakeDir(dir string){
	exists,err := PathExists(dir)
	if !exists {
		fmt.Println(err)
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil{
			fmt.Println(err)
		}
	}
}
// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadCsv()  {
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			xlog.WarningLog("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
	}
	
	
}

