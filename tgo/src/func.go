package main 

import (
	"fmt"
)

func main() {
	a,b := f1(3,"123")
	fmt.Println("main",a);
	fmt.Println("main",b);
	all("a","b")
	all("abc","123")
}

func f1(a int,b string)(a1 string,b1 string){
	fmt.Println(a)
	fmt.Println(b)
	a1 = "abc"
	b1 = "abcdefg"
	return
}
/*
* 可变长参数
*/
//其参数是一个数组
func all(str... string){
	fmt.Println(str)
	for a,b := range str{
		fmt.Println("array index:",a," value: ",b);
	}
}






