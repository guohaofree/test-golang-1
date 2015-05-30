package main

import (
	"fmt"
)

func main(){
	m := make(map[string]int64);
	m["a1"] = 77;
	fmt.Println(m);
	fmt.Println(len(m));
	m["a2"] = 11;
	fmt.Println(len(m));
	delete(m,"a1");
	fmt.Println(len(m));
	fmt.Println(m["a1"]);//被删除后的元素或不存在的元素为默认值而不是nil
	v1,v2 := m["a1"]; //如果不存在则v2为false
	fmt.Println("v1:",v1);
	fmt.Println("v2",v2);
	
	m2 := map[int32]string{27:"abc",15:"xyz"}//定义的时候直接初始化值
	fmt.Println(m2);
	fmt.Println(m2[15]);
	
	
	
	
	
}

