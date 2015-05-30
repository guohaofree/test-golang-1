package main

import (
	"io/ioutil"
	"fmt"
)

func main(){
	 // read whole the file
    b, err := ioutil.ReadFile("i:\\input.txt")
    if err != nil {
    	fmt.Print(err);
    }
	fmt.Printf("%s",b);
    // write whole the body
    err = ioutil.WriteFile("i:\\output.txt", b, 0644)
    if err != nil {
    	fmt.Print(err)
    }
    
	
}
