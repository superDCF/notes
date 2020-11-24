package main
import (
	"sync"
	"log"
)
func main() {
	var s sync.Once
	log.Println(s)
}