/*
创建时间: 2020/2/21
作者: zjy
功能介绍:

*/

package app

import (
	"bufio"
	"fmt"
	"github.com/showgo/csvdata"
	"github.com/showgo/proxy"
	"os"
	"strings"
)

//读取控制台命令
func ReadConsle()  {
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	
	defer proxy.AppWG.Done()
	for EndFlag.IsOpen() {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "reloadcsv":
			csvdata.ReLoadPublicCsvData()
		case "close":
			CloseApp()
			break
		}
	}
}
