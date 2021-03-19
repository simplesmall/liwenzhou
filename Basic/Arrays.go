package Basic

import (
	"fmt"
	"liwenzhou/Common"
)

// 数组
func Array()  {
	Common.TitleTips("数组初始化: ")
	initArrayOne()
	initArrayTwo()
	initArrayThree()
	Common.TitleTips("遍历数组: ")
	iteratorArray()
	multiDimension()
	modifyArrayTest()
	arrayPractice()
	findSumIndex()
}

/*
定义数组: var 数组变量名 [元素数量] T  eg: var arr1 [10] int
*/
func initArrayOne()  {
	Common.FuncTips("initArrayOne")
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println("testArray [3]int :",testArray)                      //[0 0 0]
	fmt.Println("numArray = [3]int{1, 2} :",numArray)                       //[1 2 0]
	fmt.Println("cityArray = [3]string{\"北京\", \"上海\", \"深圳\"} :",cityArray)                      //[北京 上海 深圳]
}

// 与方法一相比省略输入数组个数(长度)
func initArrayTwo()  {
	Common.FuncTips("initArrayTwo")
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println("var testArray [3]int :",testArray)                          //[0 0 0]
	fmt.Println("var numArray = [...]int{1, 2} :",numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println("[...]string{\"北京\", \"上海\", \"深圳\"} :",cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

// 与方法二相比使用索引下标来赋值
func initArrayThree()  {
	Common.FuncTips("initArrayThree")
	a := [...]int{1: 1, 3: 5}
	fmt.Println("Init array value by [...]int{1: 1, 3: 5}  :",a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}

// 数组遍历  迭代和range两种方法
func iteratorArray()  {
	Common.FuncTips("iteratorArray")
	var a = [...]string{"北京  ㄟ( ▔, ▔ )ㄏ", "上海 w(ﾟДﾟ)w", "深圳 ━┳━　━┳━"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历 i v可以在实际使用中选择忽略 以 _ 匿名变量处理
	for index, value := range a {
		fmt.Println(index, value)
	}
}

// 多维数组 多维数组只有第一层可以使用... 来让编译器自动推导数组长度
func multiDimension(){
	Common.FuncTips("multiDimension")
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t   ㄟ( ▔, ▔ )ㄏ ", v2)
		}
		fmt.Println()
	}
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func modifyArrayTest() {
	Common.FuncTips("modifyArrayTest")
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println("after modify , a is :",a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println("after modify , b is :",b)  //[[1 1] [1 1] [1 1]]
}
func modifyArray(x [3]int) {
	x[0] = 100
}
func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

// 练习题 : [1,3,5,7,8] 数组元素和
func arrayPractice()  {
	Common.FuncTips("arrayPractice")
	a := [...]int{1,3,5,7,8}
	sum := 0
	for _,v := range a {
		sum +=v
	}
	fmt.Println("sum of array [1,3,5,7,8] is :",sum)
}
// 练习题 找出 [1,3,5,7,8] 中值相加合为8 的下标
func findSumIndex()  {
	Common.FuncTips("findSumIndex")
	a := [...]int{1,3,5,7,8}
	target := 8
	for i1,v := range a{
		for i2, j := range a {
			if i1!=i2 && v+j == target && i1<i2{
				fmt.Println("Got you ! index is : ",i1, " and ",i2)
			}
		}
	}
}