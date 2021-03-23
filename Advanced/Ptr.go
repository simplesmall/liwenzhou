package Advanced

import (
	"fmt"
	"liwenzhou/Common"
)

func Ptr()  {
	Common.TitleTips("Go 指针")
	addType()
	getPtr()
	newPtr()
	makePtr()
}

// 指针地址和指针类型
func addType()  {
	Common.FuncTips("addType")
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
}

// 指针取值  &取出地址，*根据地址取出地址指向的值。
func getPtr()  {
	Common.FuncTips("getPtr")
	a := 10
	ap := &a
	apv := *ap
	fmt.Println("value of a is: ",a)
	fmt.Println("ptr of a is: ",ap)
	fmt.Println("after *,value of apv is: ",apv)
}

//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，
//否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
func newPtr()  {
	Common.FuncTips("newPtr")
	a :=new(int)
	b := new(bool)
	fmt.Printf("Type of a is : %T,\tptr add is %v,\tvalue is %v\n",a,a,*a)
	fmt.Printf("Type of b is : %T,\tptr add is %v,\tvalue is %v\n",b,b,*b)
}

// make 指针
// 只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，
// 而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针
func makePtr()  {
	Common.FuncTips("makePtr")
	var b map[string]int
	b = make(map[string]int, 10)
	b["咋呼的了"] = 100
	fmt.Println("make init map: ",b)
}