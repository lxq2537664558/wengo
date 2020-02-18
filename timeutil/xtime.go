/*
创建时间: 2019/11/24
作者: zjy
功能介绍:
日期相关辅助函数
*/

package timeutil

import (
	"fmt"
	"time"
)

var (
	TimeAllTemplate    string // 带毫秒的模板
	DateTemplate       string // 日期模板
	HHmmssTemplate     string // 时分秒
	YMDHTemplate       string
	FileTemplateDate   string // 日期模板
	FileTemplatemsYMDH string
)

func init() {
	TimeAllTemplate = "2006-01-02 15:04:05.000" // 常规类型
	DateTemplate = "2006-01-02"                 // 只有日期
	HHmmssTemplate = "15:04:05"                 // 时分秒
	YMDHTemplate = "2006-01-02 15:00:00"        // 常规类型
	FileTemplateDate = "20060102"
	FileTemplatemsYMDH="2006010215"
}

//获取当前时间戳秒
func GetCurrentTimeS() int64 {
	now := time.Now()
	return now.Unix()
}
//获取当前时间戳毫秒
func GetCurrentTimeMs() int64 {
	return int64(GetCurrentTimeNano() / int64(time.Millisecond))
}
//获取当前时间戳纳秒
func GetCurrentTimeNano() int64 {
	now := time.Now()
	return now.UnixNano()
}

func getTimeStr(timeNano int64,template string) string  {
	return time.Unix(0,timeNano).Format(template)
}

func GetHMSStr(timeNano int64) string {
	return getTimeStr(timeNano,HHmmssTemplate)
}

func GetTimeALLStr(timeNano int64) string {
	return getTimeStr(timeNano,TimeAllTemplate)
}


func GetYearMonthDayFromatStr(timeNano int64) string {
	nowTime := time.Unix(0,timeNano)
	datestr := fmt.Sprintf("%d%02d%02d",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day())
	return datestr
}
func GetYearMonthDayHourFromatStr(timeNano int64) string {
	nowTime := time.Unix(0,timeNano)
	datestr := fmt.Sprintf("%d%02d%02d_%2d",
		nowTime.Year(),
		nowTime.Month(),
		nowTime.Day(),
		nowTime.Hour())
	return datestr
}

func GetDateFileName(timeNano int64) string {
	return getTimeStr(timeNano,FileTemplateDate)
}