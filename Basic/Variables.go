package Basic

import (
	"fmt"
	"liwenzhou/Common"
)

// 变量
func Variables() {
	baseUse()

	batchDeclare()

	manyVarInit()

	shortVar()

	anonymousVar()
}

// 基础用法
func baseUse() {
	Common.FuncTips("baseUse")
	var tempInt int
	fmt.Println("tempInt not give value is : ", tempInt)
	// 可以省略var 定义符以及不指定变量数据类型(会自动判断和给定数据类型) == > 类型推导
	tempStr := "My dog is hot"
	fmt.Println("tempStr using := not give data type is :", tempStr)
}

// 批量声明
func batchDeclare() {
	Common.FuncTips("batchDeclare")
	var (
		patchName string
		patchAge  int
	)
	fmt.Println("batch operation value is: ", patchName, " Age is :", patchAge)
}

// 多个变量初始化
func manyVarInit() {
	Common.FuncTips("manyVarInit")
	var manyNum, manyId = "HUET6784563", 67783
	fmt.Println("Many operation result: ", manyNum, " ManyId is: ", manyId)
}

// 短变量
func shortVar() {
	Common.FuncTips("shortVar")
	n := 120
	m := "MME"
	fmt.Println("shortVar n : ", n, " m : ", m)
}

// 匿名变量
func anonymousVar() {
	Common.FuncTips("anonymousVar")
	x, y := foo()
	fmt.Println("All receive value is :", x, y)
	intx, _ := foo()
	fmt.Println("x receive value is :", intx)
	_, stry := foo()
	fmt.Println("y receive value is :", stry)
}

func foo() (int, string) {
	return 111, "MyFooFunc"
}
