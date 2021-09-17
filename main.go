package main

import "fmt"

/*	map:类似于其他语言字典，或者哈希表的东西表现为 key-value

      m    := map[int]bool{}
	变量名 := map[key的类型] value的类型{}

*/

func main() {
	fmt.Println("--------------------------")
	// m2 := make(map[string]string)

	var m map[string]string
	m = map[string]string{} //  这种 声明一定要赋值
	m["name"] = "张三"
	m["sex"] = "男生"
	fmt.Println(m) // map[name:张三 sex:男生]
	fmt.Println("string{} ↑↑↑--------------------------")

	m1 := map[string]string{}
	m1["name"] = "李四"
	m1["sex"] = "女生"
	fmt.Println(m1) // map[name:张三 sex:男生]
	fmt.Println("string{} ↑↑↑--------------------------")

	m2 := map[int]bool{}
	m2[1] = true
	m2[2] = false
	fmt.Println(m2) // map[1:true 2:false]
	fmt.Println("bool{} ↑↑↑--------------------------")

	m3 := map[int]interface{}{}
	m3[1] = 123
	m3[2] = false
	m3[3] = "strrr"
	m3[4] = []int{1, 2, 3}
	fmt.Println(m3) //	 map[1:123 2:false 3:strrr 4:[1 2 3]]
	fmt.Println("interface ↑↑↑--------------------------")

	m4 := map[int]interface{}{}
	m4[1] = 123
	m4[2] = false
	m4[3] = "strrr"
	m4[4] = []int{1, 2, 3}
	delete(m4, 4)                // 	删除第四项
	fmt.Println("删除之后", m4)      // 	map[1:123 2:false 3:strrr]
	fmt.Println("m4长度", len(m4)) // 	map[1:123 2:false 3:strrr]
	fmt.Println("delete ↑↑↑--------------------------")

	m5 := map[string]interface{}{}
	m5["a"] = 123
	m5["b"] = false
	m5["c"] = "strrr"
	m5["d"] = []int{1, 2, 3}

	for k, v := range m5 {
		fmt.Println(k, v)
	}
	fmt.Println("range 循环数组 ↑↑↑--------------------------")

	// fmt.Println("[] ↓↓↓ ********************************")
	// fmt.Println("[] ↑↑↑ ********************************")

}

// fmt.Println("[] ↓↓↓ ********************************")
// fmt.Println("[] ↑↑↑ ********************************")
