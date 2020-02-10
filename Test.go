/*
创建时间: 2019/11/24
作者: zjy
功能介绍:

*/

package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// var a int
	// a = 100
	// var c float32
	// c = 1.005
	// MyPrintf(a)
	// MyPrintf(c)
	// PathPackage()
	// filepathTest()
	// var str string
	// str = "ddasd"
	//
	// strarr :=[]string{"cc"}
	//
	// upstr := xutil.Capitalize(str)
	// fmt.Println(upstr)
	// fmt.Println(strings.Join(strarr,"|"))
	aa := "[]int"
	switch aa {
	case "[]int":
		fmt.Println("ok中国")
	}
	
	var str []string
	strarr := strings.Split("",",")
	fmt.Println(ValidArrIndex(strarr,0))
	fmt.Println(ValidArrIndex(str,10))
}



func MyPrintf(in interface{}) {
	fmt.Println(fmt.Sprintf("%v", in))
}

func PathPackage() {
	
	// 返回路径的最后一个元素
	fmt.Println(path.Base("./a/b/c"));
	// 如果路径为空字符串，返回.
	fmt.Println(path.Base("/a/d"));
	// 如果路径只有斜线，返回/
	fmt.Println(path.Base("///"));
	
	// 返回等价的最短路径
	// 1.用一个斜线替换多个斜线
	// 2.清除当前路径.
	// 3.清除内部的..和他前面的元素
	// 4.以/..开头的，变成/
	fmt.Println("Clean", path.Clean("./a/b/../"));
	
	// 返回路径最后一个元素的目录
	// 路径为空则返回.
	fmt.Println("Dir", path.Dir("./a/b/c"));
	
	// 返回路径中的扩展名
	// 如果没有点，返回空
	fmt.Println("Ext", path.Ext("./a/b/c/d.jpg"));
	
	// 判断路径是不是绝对路径
	fmt.Println("IsAbs", path.IsAbs("./a/b/c"));
	fmt.Println("IsAbs", path.IsAbs("/a/b/c"));
	
	// 连接路径，返回已经clean过的路径
	fmt.Println("join", path.Join("./a", "b/c", "../d/"));
	
	// 匹配文件名，完全匹配则返回true
	fmt.Println(path.Match("*", "a"));
	fmt.Println(path.Match("a/*/*", "a/b/c"));
	fmt.Println(path.Match("\\b", "b"));
	
	// 分割路径中的目录与文件
	fmt.Println(path.Split("./a/b/c/d.jpg"));
}

func filepathTest() {
	// 返回所给路径的绝对路径
	path, _ := filepath.Abs("./1.txt")
	fmt.Println(path)
	// 返回路径最后一个元素
	fmt.Println(filepath.Base("./1.txt"))
	// 分割路径中的目录与文件
	fmt.Println(filepath.Split("./a/b/c/d.jpg"))
	fmt.Println(filepath.Join("./a", "/b", "/c"))
	
}
