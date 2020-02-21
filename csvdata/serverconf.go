//excle生成文件请勿修改
package csvdata

import (
	"fmt"
	"github.com/showgo/csvparse"
	"github.com/showgo/xutil"
	"github.com/showgo/xlog"
	"sync/atomic"
)

var serverconfAtomic atomic.Value

type  Serverconf struct {
	Server_id int //#服务器id 字段名称  server_id
	Server_kind int //服务器类型 字段名称  server_kind
	Server_name string //服务器名称 字段名称  server_name
	Out_addr string //外部连接的地址 字段名称  out_addr
	Out_prot string //外部连接端口 字段名称  out_prot
	Max_connect int //最大连接数 字段名称  max_connect
	Send_maxsize int //发包最大数量 字段名称  send_maxsize
	Rec_maxsize int //收包最大字节 字段名称  rec_maxsize
	Write_cap_num int //连接写的包队列大小 字段名称  write_cap_num
}

func SetServerconfMapData(csvpath  string ) {
  	defer xlog.RecoverToStd()
	serverconfAtomic.Store(getServerconfUsedData(csvpath))
}

func getServerconfUsedData(csvpath  string ) map[int]*Serverconf{
    csvmapdata := csvparse.GetCsvMapData(csvpath + "/serverconf.csv")
	tem := make(map[int]*Serverconf)
	for _, filedData := range csvmapdata {
		one := new(Serverconf)
		for filedName, filedval := range filedData {
			isok := csvparse.ReflectSetField(one, filedName, filedval)
			xutil.IsError(isok)
			if _,ok := tem[one.Server_id]; ok {
				fmt.Println(one.Server_id,"重复")
			}
		}
		tem[one.Server_id] = one
	}
	return tem
}

func GetServerconfPtr(server_id int) *Serverconf{
    alldata := GetAllServerconf()
	if alldata == nil {
		return nil
	}
	if data, ok := alldata[server_id]; ok {
		return data
	}
	return nil
}

func GetAllServerconf() map[int]*Serverconf{
    val := serverconfAtomic.Load()
	if data, ok := val.(map[int]*Serverconf); ok {
		return data
	}
	return nil
}
