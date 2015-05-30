package main 

import (
	"encoding/json"
	"fmt"
)
/*
* 实验证明,转化为json的结构体中属性字段
* 首字母必须是大小,并且可以通过 field tags来标明转化为json时使用的 key名称
* 如果首字母是小写,则添加了 field tags 也不会转化这个字段
*
*同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样的一个好处是：
*当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写，
*即可轻松解决这个问题。
*/
type Struct01 struct{
	A1 string	`json:"a1"`
	a2 int
	a3 []string `json:"a3"`
}
type Struct02 struct{
	A1 string
	A2 []string `json:"a2"`//将json转化为结构体时,json中也用a2,但在结构体中调用时当然还是用A2
}
func main() {
	/*
	*将go对象转化为json元素
	*/
	v1,error := json.Marshal(5.5)
	if error==nil {
		fmt.Println("v1: ",v1,"  ",string(v1));
	}
	//----------------结构体 转化为 json
	/*
	*注意struct赋值语句,每个参数后都要加',',最后一个参数也要加
	*/
	struct01 := Struct01{
		A1:"ab",
		a2:23,
		a3:[]string{"aa","bb","ccc"},
	}
	fmt.Println(struct01);
	v2,error := json.Marshal(struct01)
	if error == nil{
		fmt.Println("v2: ",string(v2))
	}
	//--------------- json 转化为 结构体
	b1 := []byte(`{"A1":"abc","a2":["ab","cd","efg"]}`)
	var v3 Struct02
	json.Unmarshal(b1,&v3) //注意这个地址&必加,不然会报错
	fmt.Println("v3: ",v3.A2[2])
	//------------  json 转化为 map
	b2 := []byte(`{"a1":3.1415,"a2":["ab","cd","efg"]}`)
	var m1 map[string]interface{}
	json.Unmarshal(b2,&m1)
	num := m1["a1"].(float64)
	fmt.Println(num)
	
	/*
	*  encoder decoder json与io流转换
	*/
	
	
	
	
}

