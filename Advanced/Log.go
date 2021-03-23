package Advanced

import (
	"fmt"
	"liwenzhou/Common"
	"log"
	"os"
)

func Log()  {
	Common.TitleTips("log")
	//logPrint()
	logSetting()
	//logOutSetting()
}

// 最基础的log使用
func logPrint()  {
	Common.FuncTips("logPrint")
	log.Printf("This is normal log printf :%v\n",23)
	log.Fatalln("This is fatal log output")
	log.Panicln("This is panic log output") // unreachable because upper fatal is candidate
}

// 配置logger
func logSetting()  {
	Common.FuncTips("logSetting")
	log.Printf("This is normal log printf :%v\n",10010001)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Printf("After set floag log printf :%v\n",10010001)
	log.SetPrefix("[简单前缀]")
	log.Println("This is a normal log after setPrefix... ")
}

// 配置output 日志输出
func logOutSetting()  {
	Common.FuncTips("logOutSetting")
	logFile,err := os.OpenFile("./xx.log",os.O_CREATE |os.O_WRONLY | os.O_APPEND,0644)
	if err != nil {
		fmt.Println("open log file failed,err:",err)
		return
	}
	log.SetOutput(logFile)
	log.Println("This is first log after set output...")
}