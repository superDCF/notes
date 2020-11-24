package main
import (
	"log"
)
func main() {
	var a = []int{0,1,2,3,4,5}
	b:=a[:3:5]
	log.Printf("b=%v",b,len(b),cap(b))
	var aa = [...]int{0,1,2,3,4,5}
	aa = append(aa,6)
	log.Println(aa)
}