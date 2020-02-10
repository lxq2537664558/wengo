//excle生成文件请勿修改
package csvdata

import (
	"github.com/showgo/csvparse"
	"github.com/showgo/xutil"
)

var ServerconfCsv map[int]*Serverconf

type  Serverconf struct {
	ServerId int //#服务器id 字段名称  ServerId
	ServerType int //服务器类型 字段名称  ServerType
	ServerName string //服务器名称 字段名称  ServerName
	OutAddr string //外部连接的地址 字段名称  OutAddr
	OutPort string //外部连接端口 字段名称  OutPort
	MaxConnect int //最大连接数 字段名称  MaxConnect
	SendMaxsize int //发包最大数量 字段名称  SendMaxsize
	RecMaxsize int //收包最大字节 字段名称  RecMaxsize
}

func AsynSetServerconfCsvMapData() {
	 go SetServerconfCsvMapData()
}

func SetServerconfCsvMapData() {
    if ServerconfCsv == nil {
		ServerconfCsv = make(map[int]*Serverconf)
	}
	tem := getServerconfCsvUsedData("./csv/")
	ServerconfCsv  = tem
}

func getServerconfCsvUsedData(csvpath  string ) map[int]*Serverconf{
    csvmapdata := csvparse.GetCsvMapData(csvpath + "serverconf.csv")
	tem := make(map[int]*Serverconf)
	for _, filedData := range csvmapdata {
		one := new(Serverconf)
		for filedName, filedval := range filedData {
			isok := csvparse.SetFieldReflect(one, filedName, filedval)
			xutil.IsError(isok)
		}
		tem[one.ServerId] = one
	}
	return tem
}