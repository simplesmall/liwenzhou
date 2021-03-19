package Basic

import (
	"fmt"
	"liwenzhou/Common"
	"math"
	"strings"
)

// 数据类型
/*
整型、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）
*/
func DataType()  {
	hex()
	octal()
	decimal()
	floatCase()
	complexCase()
	stringCase()
	byteRune()
	changeString()
	typeTransfer()
}

/*
特殊整型: uint int 根据操作系统是多少位的变化得到,uintptr 是无符号整形,用于存放一个指针
*/

// 二进制
func hex()  {
	Common.FuncTips("hex")
	var c int  = 0xff
	fmt.Printf("oxff show hex is : %x ,show decimal is : %d\n",c,c)
}

// 八进制
func octal()  {
	Common.FuncTips("octal")
	var b int = 077
	fmt.Printf("077 八进制转换展示 : %o ,十进制展示 : %d \n",b,b)
}

// 十进制
func decimal()  {
	Common.FuncTips("decimal")
	var a int = 10
	fmt.Printf("十进制展示10 : %d \n", a)  // 10
	fmt.Printf("二进制展示10 : %b \n", a)  // 1010  占位符%b表示二进制
}

//浮点数
func floatCase()  {
	Common.FuncTips("floatCase")
	pi := math.Pi
	fmt.Printf("const float value Pi is : %f , limit decimal point : %.10F\n",pi,pi)
	fmt.Printf("science 2.3e23 show float is : %.3f\n",2.3e23)
	fmt.Printf("Max float64 is : %.2f\n\n",math.MaxFloat64)
}

// 复数
func complexCase()  {
	Common.FuncTips("complexCase")
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println("complex c1 is : ",c1, " c2 is : ",c2)
}

// 字符串
func stringCase()  {
	Common.FuncTips("stringCase")
	fmt.Println("str := \"c:\\Code\\les\tson1\\go.exe\"")
	fmt.Println("strings repeat using : ",strings.Repeat("#+",5))
	fmt.Println("string split  Split(\"wahaha\",\"a\") : ",strings.Split("wahaha","a"))
	fmt.Println("many line string : ",`
	this is line one,
		this is two`)
}

// byte and rune
/*
	uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	rune类型，代表一个 UTF-8字符。
	当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
*/
func byteRune()  {
	Common.FuncTips("byteRune")
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	// UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果
	// 104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³)
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	// 104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
	fmt.Println()
	//字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	//rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
}

// 修改字符串
/*
要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，
都会重新分配内存，并复制字节数组。
*/
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println("big change to pig using force transfer : ",string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println("白萝卜 change to 红萝卜 by rune : ",string(runeS2))
}

//类型转换
func typeTransfer()  {
	Common.FuncTips("typeTransfer")
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("float transfer to int : %d ,target type : %T, source type :  %T\n",c,c,math.Sqrt(float64(a*a + b*b)))
}