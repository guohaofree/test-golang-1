package main

import (
	"fmt"
	"os"
	"io"
)
 func main(){
// 	f,_:=os.Create("i:\\a.txt");
// 	defer f.Close();
// 	f.WriteString("1234");
// 	fmt.Println("1");
 	
 	f1,_:=os.Create("i:\\a.txt");
 	defer f1.Close();
 	reader := io.LimitReader(f1,1);
 	byte1:=make([]byte,1);
 	n,err:=reader.Read(byte1);
 	if err!=nil{ 
 		fmt.Print(err);
 	}
// 	fmt.Printf("%s",byte1);
 	fmt.Println(string(byte1[:n]));
 }
 
 
