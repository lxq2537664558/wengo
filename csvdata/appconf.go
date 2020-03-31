//excle生成文件请勿修改
package csvdata

import (
	"fmt"
	"github.com/showgo/csvparse"
	"github.com/showgo/xutil"
	"github.com/showgo/xlog"
	"sync/atomic"
)

var appconfAtomic atomic.Value

type  Appconf struct {
	App_id int //#服务器id 字段名称  app_id
	App_kind int //服务器类型 字段名称  app_kind
	App_name string //服务器名称 字段名称  app_name
	Out_addr string //外部连接的地址 字段名称  out_addr
	Out_prot string //外部连接端口 字段名称  out_prot
	Max_connect int //最大连接数 字段名称  max_connect
	Msglen_size uint8 //消息包长字节大小2 字段名称  msglen_size
	Max_msglen uint32 //消息最大长度 字段名称  max_msglen
	Write_cap_num int //连接写的包队列大小 字段名称  write_cap_num
}

func SetAppconfMapData(csvpath  string ) {
  	defer xlog.RecoverToStd()
	appconfAtomic.Store(getAppconfUsedData(csvpath))
}

func getAppconfUsedData(csvpath  string ) map[int]*Appconf{
    csvmapdata := csvparse.GetCsvMapData(csvpath + "/appconf.csv")
	tem := make(map[int]*Appconf)
	for _, filedData := range csvmapdata {
		one := new(Appconf)
		for filedName, filedval := range filedData {
			isok := csvparse.ReflectSetField(one, filedName, filedval)
			xutil.IsError(isok)
			if _,ok := tem[one.App_id]; ok {
				fmt.Println(one.App_id,"重复")
			}
		}
		tem[one.App_id] = one
	}
	return tem
}

func GetAppconfPtr(app_id int) *Appconf{
    alldata := GetAllAppconf()
	if alldata == nil {
		return nil
	}
	if data, ok := alldata[app_id]; ok {
		return data
	}
	return nil
}

func GetAllAppconf() map[int]*Appconf{
    val := appconfAtomic.Load()
	if data, ok := val.(map[int]*Appconf); ok {
		return data
	}
	return nil
}
