package Basic

import (
	"fmt"
	"liwenzhou/Common"
)

func Slice()  {
	Common.TitleTips("slice")
	defineSlice()
	lengthCap()
	cutTwice()
	totalSlice()
	makeSlice()
	emptySlice()
	copySlice()
	iteratorSlice()
	appendSlice()
	deleteSlice()
}

func defineSlice()  {
	Common.FuncTips("defineSlice")
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

// 切片长度和容量
// 为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
func lengthCap()  {
	Common.FuncTips("lengthCap")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("slice a : %v\n",a)
	fmt.Printf("s:%v len(s):%v cap(s):%v  a[:] : %v\n", s, len(s), cap(s),a[:])
}

// 对切片进行再切片
// 对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。
func cutTwice()  {
	Common.FuncTips("cutTwice")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("slice a : %v\n",a)
	fmt.Printf("s := a[1:3] : %v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2 := s[3:4] : %v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

// 完整切片表达式
// a[low:high:max] len=high-low cap=max-low
func totalSlice()  {
	Common.FuncTips("totalSlice")
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]
	fmt.Printf("slice a : %v\n",a)
	fmt.Printf("slice t (a[1:3:5]): %v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

// 使用make动态创建切片 make([]T, size, cap)
func makeSlice()  {
	Common.FuncTips("makeSlice")
	a := make([]int,2,10)
	fmt.Println("slice a : ",a," len of a : ",len(a)," cap of a : ",cap(a))
}

// 判断slice是否为空要用len,不能用nil 切片不能直接比较
func emptySlice()  {
	Common.FuncTips("emptySlice")
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println("slice s1: ",s1," slice s2: ",s2==nil,"slice s3: ",s3)
}

// 切片的赋值拷贝  前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意
func copySlice()  {
	Common.FuncTips("copySlice")
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println("slice s1: ",s1," slice s2: ",s2) //[100 0 0] [100 0 0]
}

// 切片遍历
func iteratorSlice()  {
	Common.FuncTips("iteratorSlice")
	s := []int{1, 3, 5}

	fmt.Println("使用fori 方法遍历切片,输出下标及数值:")
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	fmt.Println("使用for range 方法遍历切片,输出下标及数值:")
	for index, value := range s {
		fmt.Println(index, value)
	}
}

// slice append
func appendSlice()  {
	Common.FuncTips("appendSlice")
	var s []int
	fmt.Println("original slice s is: ",s)
	s1 := append(s,1)
	fmt.Println("s append 1 slice s1 is: ",s1)
	temps :=  []int{2,3,4}
	s2 := append(s1,temps...)
	fmt.Println("s append 2,3,4 slice s2 is: ",s2)
}

// slice copy,切片的复制,复制之后的切片与源切片相隔离,直接赋值 b:=a则由于切片是应用类型,所以改变会同步更改
// copy(destSlice,srcSlice,[]T)

// 删除slice中的元素
// 切片中没有自带删除元素,但可以通过索引操作实现删除指定索引下标的元素
// 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
func deleteSlice()  {
	Common.FuncTips("deleteSlice")
	a := []int{20,21,22,23,24,25}
	fmt.Println("slice a :",a)
	a = append(a[:2],a[3:]...)
	fmt.Println("After delete index 2,slice is :",a)
}