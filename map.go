package main

import "fmt"

func main() {
	//Khai báo map
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	if myMap == nil {
		fmt.Println("myMap = nil")
	} else {
		fmt.Println("!= nil")
	}
	fmt.Println()
	var myMap1 map[string]int
	fmt.Println(myMap1)

	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	}

	//Khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}
	fmt.Println(myMap2)

	// them 1 phan tu vao map
	myMap2["key5"] = 5
	myMap2["key6"] = 6
	fmt.Println(myMap2)

	//delete 1 phan tu trong map : delete(map, key)
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	//length map
	fmt.Println(len(myMap2))

	//Map là một reference type
	myMap3 := myMap2
	fmt.Println(myMap3)

	myMap3["key5"] = 9
	fmt.Println(myMap2)

	// truy cập 1 phẩn tử trong map
	value, found := myMap2["key1000"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("khong tim thay")
	}
	// trong map không có toán tử ==
	if myMap2 == myMap3 {
		fmt.Println(value)
	}
}
