package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	local := time.FixedZone("", -28800)
	ss := time.Now().In(local).Format("20060102")
	log.Println(ss)
	m := "10001"
	log.Println(strconv.Atoi(m))
	log.Println(time.Now().UTC().Hour())
	log.Println(time.Now().Zone())
	loc, _ := time.LoadLocation("America/New_York")
	printTime(time.Now().In(loc))
	log.Println(time.Now().UTC().Hour(), time.Now().UTC().Minute())

}
func printTime(t time.Time) {
	zone, offset := t.Zone()
	fmt.Println(t.Format(time.Kitchen), "Zone:", zone, "Offset UTC:", offset)
}
