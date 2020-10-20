package main
 import (
     "fmt"
     "time"
     "sync"

 )

 func main() {
    var wg sync.WaitGroup
    wg.Add(1) 
  go func() {
      count("sheep")
      wg.Done()
 }()
   func count(thing string, c chan string) {
     for i := 1; i <=5;i++ {
     c <- thing
       fmt.Println(i, thing)
       time.Sleep(time.Millisecond * 500)
     }
   }