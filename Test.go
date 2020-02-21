/*
创建时间: 2019/11/24
作者: zjy
功能介绍:

*/

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
	"github.com/showgo/csvdata"
	"github.com/showgo/protobuf/pb/login_proto"
	"github.com/showgo/xutil"
	"math"
	
	// "github.com/jinzhu/gorm"
	"github.com/showgo/dbutil"
	"sync"
)

type User struct{
	Name string
	Age int
}

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
	
	// user := new(User)
	
	fmt.Println(xutil.IsLittleEndian() ,math.MaxInt32)
}

func TestSwitch(aname int)  {
	switch aname {
	case 1,10:
		fmt.Println("aaa")
	}
}

func DbTest()  {

	csvdata.SetDbconfMapData("./csv")
		// gorm.Open()
	gamedb := dbutil.OpenDB(csvdata.GetDbconfPtr("gamedb"))
	if gamedb == nil {
		return
	}
	fmt.Println(dbutil.CheckTableExists(gamedb,"gamedb","Account"))
}


func  protobufTest(){
	person := &login_proto.Person{
		Id:10,
		Name:"郑蛟元",
	}
	onePhone := new(login_proto.Phone)
	onePhone.Type = 2
	onePhone.Number = "15223153231"
	onePhone1 := new(login_proto.Phone)
	onePhone1.Type = 3
	onePhone1.Number = "733528"
	person.Phones = append(	person.Phones,onePhone,onePhone1)
	
	fmt.Println(person)
	data,erro := proto.Marshal(person)
	if erro  != nil {
		fmt.Println(erro)
	}
	fmt.Println(data,len(data))
	person2 := new(login_proto.Person)
	
	proto.Unmarshal(data,person2)
	fmt.Println("person2",person2)
}
