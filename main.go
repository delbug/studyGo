package main

import (
	"fmt"
)

/*
	函数：func
	函数首字母大写可以被别的包调用，小写不可以被别的包调用
*/

func main() {
	fmt.Println("[函数] ↓↓↓ ********************************")
	a(0, "str字符串0") // data1: 0 ,data2: str字符串0
	a(2, "str字符串2") // data1: 2 ,data2: str字符串2
	fmt.Println("[函数] ↑↑↑ ********************************")

	a, b := swap("Google", "Runoob")
	fmt.Println(a, b) // Runoob Google
	fmt.Println("[] ↑↑↑ ********************************")

	c, d := funcNum("1", "strrrr")
	fmt.Println(c, d) // 1 等于: 1
	fmt.Println("[funcNum] ↑↑↑ ********************************")

	e, f := funcStr(1, 2)
	fmt.Println(e, f) // 2  !== 1
	fmt.Println("[funcStr] ↑↑↑ ********************************")

	g, h := funcName("张三", "李四")
	fmt.Println(g, h) // 张三1 李四2
	fmt.Println("[funcName] ↑↑↑ ********************************")

	niming := func(data string) {
		fmt.Println(data)
	}
	niming("是匿名函数-") // 是匿名函数-
	fmt.Println("[我是匿名函数] ^^^ ****************************")

	budingxiang(1, "2", "3", "4", "5") // 1 [2 3 4 5]
	// 类似于 arguments
	fmt.Println("[不定向参数] ^^^ ****************************")

}

func a(data1 int, data2 string) {
	fmt.Println("data1:", data1, ",data2:", data2)
	if data1 > 1 {
		fmt.Println(data1)
	} else {
		fmt.Println(data2)
	}
}

func swap(x, y string) (string, string) {
	return y, x
}

func funcNum(x string, y string) (string, string) {
	if x == "1" {
		return x, "等于: " + x
	} else {
		return y, "不等于1,为: " + y
	}
}

func funcStr(x int, y int) (int, string) {
	if x > 1 {
		return x, ">1"
	} else {
		return y, " !== 1"
	}
}

func funcName(x string, y string) (a string, b string) {
	a = x + "1"
	b = y + "2"
	return
}

// 不定向参数
func budingxiang(x int, y ...string) {
	fmt.Println(x, y)
}

// fmt.Println("[] ↓↓↓ ********************************")
// fmt.Println("[] ↑↑↑ ********************************")
