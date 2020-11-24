package main
import (
    "sync"
	"time"
	"log"
)
func main() {
    var mu sync.Mutex
    go func() {
		mu.Lock()
		log.Println("444")
	   time.Sleep(3 * time.Second)
	   log.Println("222")
	//    mu.Unlock()
	   log.Println("333")
   }()
   time.Sleep(time.Second)
   mu.Unlock()
   log.Println("111")
   time.Sleep(3 * time.Second)
//    select {}
   log.Println("last")
}