package main

import "fmt"

func main() {
	var map1 = make(map[string]string, 3)
	map1 = map[string]string{
		"one": "1",
		"two": "2",
	}
	map1 = map[string]string{
		"three": "3",
		"four":  "4",
	}
	fmt.Println(map1)
	//fmt.Println("_cut1:", cut1, cap(cut1))
	//fmt.Println("_cut2:", cut2, cap(cut2))
	//fmt.Println("_cut3:", cut3, cap(cut3))
}
