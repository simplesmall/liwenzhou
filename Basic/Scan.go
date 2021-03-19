package Basic

import (
	"bufio"
	"fmt"
	"liwenzhou/Common"
	"os"
	"strings"
)

// scan 测试时下面函数逐个单独运行测试,避免意外退出
func Scan() {
	Common.TitleTips("Scan")
	//fmtScan()
	//scanf()
	bufioDemo()
}

// 从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
// Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
func fmtScan() {
	Common.FuncTips("fmtScan")
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println("Please input name(string) age(int) married(bool),seperation with ' ':")
	_, _ = fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。
func scanf() {
	Common.FuncTips("scanf")
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println("Please input 1:name(string) 2:age(int) 3:married(bool),seperation with ' ':")
	_, _ = fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func bufioDemo() {
	Common.FuncTips("bufioDemo")
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("读取到的输入内容: %#v\n", text)
}

//Fscan系列  Fscan  Fscanln  Fscanf
//这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
//只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。

//Sscan系列	Sscan  Sscanln  Sscanf
//这几个函数功能分别类似于fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，
//只不过它们不是从标准输入中读取数据而是从指定字符串中读取数据。
