package Common

import (
	"fmt"
	"strings"
)

// 多行分隔符
func ManyLines(n int)  {
	fmt.Println(strings.Repeat("\n",n))
}
