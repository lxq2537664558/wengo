/*
创建时间: 2020/2/2
作者: zjy
功能介绍:

*/

package apploginsv

import (
     "github.com/showgo/xengine"
)



type LoginServerFactory struct {

}



func (lsf *LoginServerFactory)CreateAppBehavor() xengine.ServerBehavior {
     ls := new(LogionServer)
     ls.OnStart()
     return ls
}

