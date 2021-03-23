package Advanced

import (
	"errors"
	"fmt"
	"liwenzhou/Common"
)

func Func() {
	Common.TitleTips("Func基础用法")
	manyParameters(0, 2, 3, 6)
	a, b := manyReturns()
	fmt.Println("Function return many values: ", a, b)
	sum, sub := nameResult()
	fmt.Println("Name result returns: ", sum, sub)
	funcVariable()
	addRes := funcAsPar(10, 20, add)
	fmt.Println("10 + 20 =: ", addRes)
	res, _ := funAsRes("+")
	fmt.Println("res(12,23),same as add(12,23): ", res(12, 23))
	anonymousFunc()
}

// 函数传入多个参数
func manyParameters(base int, x ...int) {
	Common.FuncTips("manyParameters")
	sum := base
	for _, v := range x {
		sum += v
	}
	fmt.Println("sum of all x is: ", sum)
}

// 函数多个返回值
func manyReturns() (x, y int) {
	Common.FuncTips("manyReturns")
	return 12, 44
}

// 返回值命名
func nameResult() (sum, sub int) {
	Common.FuncTips("nameResult")
	a, b := 11, 22
	sum = a + b
	sub = a - b
	return
}

type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数类型变量
func funcVariable() {
	Common.FuncTips("funcVariable")
	var c calculation
	c = add
	fmt.Println("13 + 24 =: ", c(13, 24))
	f := sub
	fmt.Println("56 - 12 =: ", f(56, 12))
}

// 高阶函数之函数作为参数
func funcAsPar(x, y int, op func(int, int) int) int {
	Common.FuncTips("funcAsPar")
	return op(x, y)
}

// 高阶函数之函数作为返回值
func funAsRes(s string) (func(int, int) int, error) {
	Common.FuncTips("funAsRes")
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 匿名函数
func anonymousFunc() {
	Common.FuncTips("anonymousFunc")
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println("将匿名函数保存到变量,x+y=: ", x+y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println("自执行函数：匿名函数定义完加()直接执行,x+y=: ", x+y)
	}(10, 20)
}

//闭包 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
func closure() {
	Common.FuncTips("closure")

}
