/*
创建时间: 2020/2/2
作者: zjy
功能介绍:

*/

package appclient

import (
     "github.com/showgo/xengine"
)



type ClientFactory struct {

}



func (this *ClientFactory)CreateAppBehavor() xengine.ServerBehavior {
     ls := new(AppClient)
     ls.OnStart()
     return ls
}

