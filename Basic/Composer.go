package Basic

import "liwenzhou/Common"

// 变量集合器
func VariablesCompose()  {
	Common.GroupTips("基础之变量和常量-变量")
	Common.TitleTips("变量")
	Variables()
}
// 常量集合器
func ConstVariableCompose()  {
	Common.TitleTips("常量")
	ConstVariables()
}

// 运算符集合器
func OperatorCompose()  {
	Common.GroupTips("运算符")
	Common.TitleTips("运算符")
	Operator()
}
// 流程控制集合器
// if-else 流程控制
func ProcessControlCompose()  {
	Common.GroupTips("流程控制")
	Common.TitleTips("if-else")
	IfElse()
}

// for循环流程控制
func ForCompose()  {
	Common.TitleTips("for循环")
	ForLoop()
}
