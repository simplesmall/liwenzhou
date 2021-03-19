package Basic

import (
	"errors"
	"fmt"
	"liwenzhou/Common"
	"os"
)

// fmt 常用函数
func Fmt() {
	Common.TitleTips("Print")
	print()
	Common.TitleTips("Fprint")
	//fprint()
	Common.TitleTips("Sprint")
	sprint()
	Common.TitleTips("Errorf")
	errorf()
	Common.TitleTips("格式化占位符")
	normalPlaceholder()
	intPlaceholder()
	floatPlaceholder()
	charsetPlaceholder()
	ptrPlaceholder()
	widthPlaceholder()
	otherPlaceholder()
}

func print() {
	Common.FuncTips("print")
	fmt.Print("在终端打印该信息。")
	name := "简单的小黑"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}

// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
func fprint() {
	Common.FuncTips("fprint")
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "稳定的大黑牛"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

// Sprint系列函数会把传入的数据生成并返回一个字符串。
func sprint() {
	Common.FuncTips("sprint")
	s1 := fmt.Sprint("给你一首歌")
	name := "温柔的小羊"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("来了一只蛤蟆")
	fmt.Println(s1, s2, s3)
}

func errorf() {
	Common.FuncTips("errorf")
	e := errors.New("假装出了一个不得了的错e")
	w := fmt.Errorf("Wrap 了一个错误:  %w", e)
	fmt.Println(w)
}

// 通用占位符
func normalPlaceholder() {
	Common.FuncTips("normalPlaceholder")
	fmt.Printf("using %%v : %v\n", 100)
	fmt.Printf("using %%v : %v\n", false)
	fmt.Printf("using %%t : %t\n", true)
	o := struct{ name string }{"小王子"}
	fmt.Printf("using %%v : %v\n", o)
	fmt.Printf("using %%#v : %#v\n", o)
	fmt.Printf("using %%T : %T\n", o)
	fmt.Printf("using %%%% : 100%%\n")
}

// 整型占位符
func intPlaceholder() {
	Common.FuncTips("intPlaceholder")
	n := 65
	fmt.Printf("show 65 using %%b : %b\n", n) // 二进制			1000001
	fmt.Printf("show 65 using %%c : %c\n", n) // 该值对应unicode值	A
	fmt.Printf("show 65 using %%d : %d\n", n) // 十进制			65
	fmt.Printf("show 65 using %%o : %o\n", n) // 八进制			101
	fmt.Printf("show 65 using %%x : %x\n", n) // 十六进制 使用a~f	41
	fmt.Printf("show 65 using %%X : %X\n", n) // 十六进制 使用A~F	41
	//%U	表示为Unicode格式：U+1234，等价于”U+%04X”
	//%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
}

// 浮点数与复数占位符
func floatPlaceholder() {
	Common.FuncTips("floatPlaceholder")
	f := 12.34
	fmt.Printf("show 12.34 using %%b : %b\n", f) // 无小数部分、二进制指数的科学计数法					6946802425218990p-49
	fmt.Printf("show 12.34 using %%e : %e\n", f) // 科学计数法，如-1234.456e+78						1.234000E+01
	fmt.Printf("show 12.34 using %%E : %E\n", f) // 	科学计数法，如-1234.456E+78					1.234000E+01
	fmt.Printf("show 12.34 using %%f : %f\n", f) // 	有小数部分但无指数部分，如123.456				12.340000
	fmt.Printf("show 12.34 using %%g : %g\n", f) // 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）	12.34
	fmt.Printf("show 12.34 using %%G : %G\n", f) // 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）	12.34
}

// 字符串和[]byte
func charsetPlaceholder() {
	Common.FuncTips("charsetPlaceholder")
	s := "一头水牛"
	fmt.Printf("show 一头水牛 using %%s : %s\n", s) // 直接输出字符串或者[]byte										一头水牛
	fmt.Printf("show 一头水牛 using %%q : %q\n", s) // 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示	"一头水牛"
	fmt.Printf("show 一头水牛 using %%x : %x\n", s) // 	每个字节用两字符十六进制数表示（使用a-f)						e4b880e5a4b4e6b0b4e7899b
	fmt.Printf("show 一头水牛 using %%X : %X\n", s) // 	每个字节用两字符十六进制数表示（使用A-F）						E4B880E5A4B4E6B0B4E7899B
}

// 指针占位符
func ptrPlaceholder() {
	Common.FuncTips("ptrPlaceholder")
	a := 10
	fmt.Printf("show 10 using %%p  : %p\n", &a)  //表示为十六进制，并加上前导的0x 0xc00000a1b0
	fmt.Printf("show 10 using %%#p : %#p\n", &a) // c00000a1b0
}

// 宽度标识符
func widthPlaceholder() {
	Common.FuncTips("widthPlaceholder")
	n := 12.34
	fmt.Printf("show 12.34 with %%f   : %f\n", n)    // 默认宽度，默认精度	12.340000
	fmt.Printf("show 12.34 with %%9f  : %9f\n", n)   // 宽度9，默认精度		12.340000
	fmt.Printf("show 12.34 with %%.2f : %.2f\n", n)  // 默认宽度，精度2		12.34
	fmt.Printf("show 12.34 with %%9.3f: %9.2f\n", n) // 宽度9，精度2		 	12.34
	fmt.Printf("show 12.34 with %%9.f : %9.f\n", n)  // 宽度9，精度0				12
}

// 其它占位符
/*
	’+’	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	’ ‘	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
	’-’	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	’#’	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
	‘0’	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
*/
func otherPlaceholder()  {
	Common.FuncTips("otherPlaceholder")
	s := "百合花"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)
}