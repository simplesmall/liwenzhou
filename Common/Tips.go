package Common

import (
	"fmt"
	"strings"
)

// 提示
func MainTips(msgType int,msg string)  {
	var tempMsgType = ""
	switch msgType {
	case 1:
		tempMsgType = "分组 :"
	case 2:
		tempMsgType = "标题 :"
	case 4:
		tempMsgType = "函数 :"
	default:
		tempMsgType = "Normal :"
	}
	fmt.Println(tempMsgType, strings.Repeat("==",8-msgType),">",msg)
}

func GroupTips(msg string)  {
	MainTips(1,msg)
}
func TitleTips(msg string)  {
	MainTips(2,msg)
}
func FuncTips(msg string)  {
	MainTips(4,msg)
}
