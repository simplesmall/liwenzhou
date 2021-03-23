package Advanced

import (
	"fmt"
	"liwenzhou/Common"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// map
func Map() {
	Common.TitleTips("Map")
	defineMap()
	traverseMap()
	traverseMapCustom()
	mapItem()
	mapSlice()
	wordsCounter()
	mapTest()
}

// define 定义map  make(map[KeyType]ValueType, [cap])
func defineMap() {
	Common.FuncTips("defineMap")
	sourceMap := make(map[string]int, 0)
	sourceMap["first"] = 111
	fmt.Println("map sourceMap is: ", sourceMap)
	fmt.Printf("Type of sourceMap is: %T,\tvalue of source[first] is %v\n", sourceMap, sourceMap["first"])
	// check if the record is exist
	v, ok := sourceMap["second"]
	if ok {
		fmt.Println("v:", v)
	} else {
		fmt.Println("second not in sourceMap")
	}
}

// 遍历map
func traverseMap() {
	Common.FuncTips("traverseMap")
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	delete(scoreMap, "小明") //将小明:100从map中删除
	fmt.Println("After delete 小明, sourceMap is: ")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

// 按照指定顺序遍历map
func traverseMapCustom() {
	Common.FuncTips("traverseMapCustom")
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// 切片中的元素为map类型时的操作
func mapItem() {
	Common.FuncTips("mapItem")
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "颜如玉"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "广寒宫"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// map中值为切片类型的操作
func mapSlice() {
	Common.FuncTips("mapSlice")
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

// 统计一个字符串中各个单词出现的次数
// "how do you do this time for you do this"
// do:3 for:1 how:1 this:2 time:1 you:2
func wordsCounter() {
	Common.FuncTips("wordsCounter")
	source := "how do you do this time for you do this"
	// 输出原string
	fmt.Println("source string: ", source)
	// 分割string存为数组
	var items []string
	items = strings.Split(source, " ")
	fmt.Println("after split,items is: ", items)
	counter := make(map[string]int, len(items))
	// 遍历计数
	for _, value := range items {
		_, ok := counter[value]
		// 原本不存在则设置为1,原本已经存在则+1
		if !ok {
			counter[value] = 1
		} else {
			counter[value] += 1
		}
	}
	fmt.Println(counter)
}

// 将map存放在其它地方之后删除本map数据不影响已存储的内容
func mapTest() {
	Common.FuncTips("mapTest")
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("source map: %+v\n", s)
	m["moto"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("delete index 1: %+v\n", s)
	fmt.Printf("map[string][] : %+v\n", m["moto"])
}
