package main

import (

	// "crypto/rand"
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)



func main() {
	// random number generator
	// fmt.Println("Hello, World!")
	// fmt.Println(rand.Intn(100))

	// random from crypto
	num, err := rand.Int(rand.Reader, big.NewInt(5))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num)
}
