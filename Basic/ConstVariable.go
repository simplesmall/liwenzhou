package Basic

import (
	"fmt"
	"liwenzhou/Common"
)

// 常量
func ConstVariables() {
	baseUseConst()
	manyDeclare()
	iotaNormal()
	iotaJump()
	iotaJumpLine()
	iotaForUnit()
	manyIotaInOneLine()
}

func baseUseConst() {
	Common.FuncTips("baseUse")
	const pi = 3.1415926
	const e = 2.7182
	fmt.Println("const pi and e is:", pi, "\t", e)
}

func manyDeclare() {
	Common.FuncTips("manyDeclare")
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println("declare many value and assign first: ", n1, n2, n3)
}

func iotaNormal() {
	Common.FuncTips("iotaNormal")
	const (
		n1 = iota //0
		n2        //1
		n3        //2
	)
	fmt.Println("iota normal using :", n1, n2, n3)
}

// 忽略某个按顺序该取到的值
func iotaJump() {
	Common.FuncTips("iotaJump")
	const (
		n1 = iota
		_
		n2
	)
	fmt.Println("iota jump value :", n1, n2)
}

// iota跳过或指定某个特定值
func iotaJumpLine() {
	Common.FuncTips("iotaJumpLine")
	const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0
	fmt.Println("iota jump in line :", n1, n2, n3, n4, n5)
}

//iota用于单位转换
func iotaForUnit()  {
	Common.FuncTips("iotaForUnit")
	const (
		_  = iota				// 弃用 0
		KB = 1 << (10 * iota)	// iota = 1
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println("flexible using iota for unit :",KB,MB,GB)
}

// 多个iota定义在一行内
func manyIotaInOneLine()  {
	Common.FuncTips("manyIotaInOneLine")
	const (
		a, b = iota + 1, iota + 3 //1,3
		c, d                      //2,4
		e, f                      //3,5
	)
	fmt.Println("declare many iota variable in one line :",a,b,c,d,e,f)
}
