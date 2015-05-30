package main

import (
	"fmt"
	
	"time"
)
func main() {
      ch := make( chan int,100 )
       go func () {
             for i := 0; i < 10; i++ {
             	time.Sleep(time.Second);
                ch <- i
            }
      }()
      
       for w := range ch {
            fmt.Println( "fmt print" , w)
//             if w > 5 {
//                   //break // 在这里break循环也可以
//                   close(ch)
//            }
      }
      ch <- 100;
      fmt.Println( "after range or close ch!" )
}