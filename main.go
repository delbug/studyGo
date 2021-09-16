package main

import (
	"fmt"
	dog "studyGo/dog"
	testPK "studyGo/testpackage"
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

}
