/*
创建时间: 2019/11/6
作者: zjy
功能介绍:
工具包
*/

package xutil

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)


func MakeDirAll(dir string) bool{
	if  StringIsNil(dir) { // 路径为nil不能创建
		return  false
	}
	exists,err := PathExists(dir)
	if !exists {
		if err != nil {
			fmt.Println(dir," 不存在需要创建",err)
		}
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil{
			fmt.Println(err)
			return false
		}
	}
	return  true
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

//获取目录
func ReadDir(path string) (*os.File, error)  {
	return  os.OpenFile(path, os.O_RDONLY, os.ModeDir)
}

//验证内置类型数组
func ValidArrIndex(arr interface{},index int) bool  {
	if arr == nil {
		return false
	}
	// 下标为负
	if   index < 0  {
		return  false
	}
	switch val := arr.(type) {
	case []int:
		return index < len(val)
	case []string:
		return index < len(val)
	case []float32:
		return index < len(val)
	default:
		fmt.Println(arr,"is an unknown type. ")
		return false
	}
	return true
}

//是否错误，有错返回 true无错返回false
func IsError(err error) bool  {
	if  err != nil {
		buf := make([]byte, 4096)
		l := runtime.Stack(buf, false)
		fmt.Printf("%v \n%s", err, buf[:l])
		return true
	}
	return  false
}

//是否错误，有错返回 true无错返回false
func IsErrorNoPrintf(err error) bool  {
	return err != nil
}


//判断字符串是否有数据  无数据返回true
func StringIsNil(str string) bool  {
	return   len(str) == 0 || str == ""
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	if StringIsNil(str) {
		return str
	}
	var upperStr string
	vv := []rune(str)
	if vv[0] >= 97 && vv[0] <= 122 {  // 后文有介绍
		vv[0] -= 32 // string的码表相差32位
		upperStr = string(vv[0]) + string(vv[1:len(vv)])
	} else {
		fmt.Println("Not begins with lowercase letter,")
		return str
	}
	
	return upperStr
}

//是否是xlsx 文件
func IsXlsx(fileName string) bool  {
	return path.Ext(fileName) == ".xlsx" && !strings.HasPrefix(fileName, "~$")
}
//验证csv行数据是否有效
//除第三行外,行没有注释 str首字符 != #  ASCII表 35
//并且id不为nil
func ValidCsvRow(str string,rownum int) bool  {
	if StringIsNil(str)  {
		return false
	}
	if rownum != 2 && str[0] == 35 {
		return  false
	}
	return true
}


//计算内存
func MemConsumed()  uint64 {
	runtime.GC() //GC，排除对象影响
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)
	return memStat.Sys
}
//获取包的字符串名称
func GetPackageStr(pkgname string)  string {
	return fmt.Sprintf("\"%s\"",pkgname)
}

//主机是否是小端序列编码
func IsLittleEndian() bool {
	n := 0x1234
	//转换获取小端的数值
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
}


func StrToUint8(str string) uint8{
	return uint8(StrToInt(str))
}

func StrToUint16(str string) uint16{
	return uint16(StrToInt(str))
}

func StrToUint32(str string) uint32{
	return uint32(StrToInt(str))
}
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

