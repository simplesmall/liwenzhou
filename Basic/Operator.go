package Basic

import (
	"fmt"
	"liwenzhou/Common"
)

// 运算符

/*
	算术运算符
	关系运算符
	逻辑运算符
	位运算符
	赋值运算符
*/
func Operator() {
	calculateOpe()
	relationOpe()
	logicOpe()
	bitOpe()
	assignOpe()
}

// 算数运算符
func calculateOpe() {
	Common.FuncTips("calculateOpe")
	counter := 1
	sum := 10
	counter++
	sum += counter
	//  2 12 1
	fmt.Println("base ++ and += :", counter, sum, 5%2)
}

// 关系运算符
func relationOpe() {
	Common.FuncTips("relationOpe")
	// true false false true false
	fmt.Println("relation result(true or false)  1 == 1,1!=1,1>2,1<2,'w'<='e':", 1 == 1, 1 != 1, 1 > 2, 1 < 2, 'w' <= 'e')
}

// 逻辑运算符
func logicOpe() {
	Common.FuncTips("logicOpe")
	a, b := true, false
	// false true false
	fmt.Println("logic result(a&&b,a||b,!a) :", a && b, a || b, !a)
}

// 位运算符
func bitOpe() {
	Common.FuncTips("bitOpe")
	// 4 8 1 15 0 -9
	fmt.Println("bit operator result(1<<2,10^2,1|0,7|12,3&8,^8) :", 1<<2, 10^2, 1|0, 7|12, 3&8, ^8)
}

//赋值运算符
func assignOpe() {
	Common.FuncTips("assignOpe")
	var temp int = 0
	temp++
	fmt.Println("temp begin with 0 ,++ is:", temp) // 1
	temp += 3
	fmt.Println("+=3 is:", temp) // 4
	temp -= 2
	fmt.Println("-=2 is:", temp) // 2
	temp *= 3
	fmt.Println("*=3 is:", temp) // 6
	temp /= 2
	fmt.Println("/=2 is:", temp) // 3
	temp %= 2
	fmt.Println("%=2 is:", temp) // 1
	temp &= 2
	fmt.Println("&=2 is:", temp) // 0
	temp |= 5
	fmt.Println("|=5 is:", temp) // 5
	temp ^= 2
	fmt.Println("^=2 is:", temp) // 7
	temp <<= 2
	fmt.Println("<<=2 is:", temp) //	28
	temp >>= 3
	fmt.Println(">>=3 is:", temp) // 3

}
