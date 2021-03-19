package Basic

import "liwenzhou/Common"

// 数据类型集合器
func DataTypeCompose()  {
	Common.GroupTips("基本数据类型")
	Common.TitleTips("数据类型")
	DataType()
}

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
func ProcessControlCompose()  {
	Common.GroupTips("流程控制")
	ProcessControl()
}
// 数组集合器
func ArrayCompose()  {
	Common.GroupTips("数组")
	Array()
}
// Fmt集合器
func FmtCompose()  {
	Common.GroupTips("Fmt")
	Fmt()
}
// Scan集合器
func ScanCompose()  {
	Common.GroupTips("Scan")
	Scan()
}