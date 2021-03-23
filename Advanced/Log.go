package Advanced

import (
	"fmt"
	"liwenzhou/Common"
)

func Log()  {
	Common.TitleTips("log")
	logPrint()
}

func logPrint()  {
	Common.FuncTips("logPrint")
	fmt.Println("log print level test")
}