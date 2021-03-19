package Basic

import (
	"fmt"
	"liwenzhou/Common"
)

// 流程控制
func ProcessControl()  {
	Common.TitleTips("if-else")
	IfElse()
	Common.TitleTips("for-loop")
	ForLoop()
	Common.TitleTips("switch")
	switchCase()
}

func IfElse()  {
	Common.FuncTips("IfElse")
	if score:=96;score>=90 {
		fmt.Println("A")
	}else if score >80{
		fmt.Println("B")
	}else {
		fmt.Println("C")
	}
}

func ForLoop()  {
	Common.FuncTips("ForLoop")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

/*
一个分支可以有多个值，多个case值中间使用英文逗号分隔
分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。
fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
*/
func switchCase()  {
	Common.FuncTips("switchCase")
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}