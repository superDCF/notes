package main

import (
	"log"
	"math/rand"
)

func main() {

	rand.Seed(121)
	log.Printf("int63=%v", rand.Int63())
}
