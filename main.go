package main

import (
	"fmt"
	dog "studyGo/study1-2/dog"
	testPK "studyGo/study1-2/testpackage"
	//
	/****************************************************************
	二
		testPK "studyGo/testpackage"
		testPK  别名  调用A: fmt.Println(testPK.A)
		***************************************************************
		testPK "studyGo/testpackage"
		testPK  别名  调用A: fmt.Println(testPK.A)
		***************************************************************
		. "studyGo/testpackage"
		调用A: fmt.Println(A)
		***************************************************************
		导出的变量或方法大写字母 开头才可以拿给别的包使用，小写开头的为私有的
		***************************************************************
		导入多个不同的包，多个包里面的变量名不能重复，不然会提示红波浪线
		***************************************************************

	三
		1
	*/)

func main() {
	fmt.Println("1-2课 ↓↓↓ ********************************")
	var a string = "main-aaaaa"
	// 关键字 变量名 变量类型 = 变量值
	c := 123
	d := true
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(testPK.A)
	fmt.Println(testPK.B)
	fmt.Println(testPK.E)
	fmt.Println(testPK.F)
	// fmt.Println(testPK.f) // 导出的变量或方法大写字母开头才可以拿给别的方法使用，小写开头的为私有的
	fmt.Println(dog.DogName)
	fmt.Println(testPK.Catname)
	fmt.Println("1-2课 ↑↑↑ ********************************")

	fmt.Println("3课 ↓↓↓ ********************************")
	fmt.Println("取值范围 ↓↓↓ ********************************")
	var num1 uint = 999 // 正整数
	var num2 int = -999 // 全部整数
	var num3 int8 = 100 // -128 --> 128 之间

	fmt.Println("uint-正整数:", num1)
	fmt.Println("int-全部整数:", num2)
	fmt.Println("int8 -128-->128:", num3)

	fmt.Println("取值范围 ↑↑↑  ********************************")

	fmt.Println("浮点型 ↓↓↓ ********************************")
	var num4 float64 = 3.1415926 // 小数点后面 2的多少次方
	fmt.Println("float64:", num4)
	fmt.Println("浮点型 ↑↑↑  ********************************")

	fmt.Println("字符串类型 ↓↓↓ ********************************")
	var str5 string = "我是字符串str--"
	fmt.Println("string:", str5)
	// fmt.Printf(format:"%T", str5)
	// fmt.Printf("str5 检测类型为:%T ", str5)

	fmt.Println("字符串类型 ↑↑↑  ********************************")

	fmt.Println("布尔值  ↓↓↓ ********************************")
	var bool6 bool = true

	fmt.Println("bool:", bool6)
	// fmt.Printf("bool6 检测类型为:%T", bool6)

	fmt.Println("布尔值  ↑↑↑  ********************************")

	fmt.Println("检测类型  ↓↓↓ ********************************")
	var str8 string = "检测类型str--"
	var bool9 bool = true
	fmt.Printf("str8 检测类型为: %T ", str8)
	fmt.Println("")
	fmt.Printf("bool9 检测类型为: %T ", bool9)
	fmt.Println("")
	fmt.Println("检测类型  ↑↑↑  ********************************")

	fmt.Println("3课 ↑↑↑ ********************************")

	// fmt.Println("  ↓↓↓ ********************************")
	// fmt.Println("  ↑↑↑  ********************************")
}
