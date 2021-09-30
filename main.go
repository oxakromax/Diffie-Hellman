package main

import (
	"math/big"
	"math/rand"
	"time"
)

func getRandomN(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
func isPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

func pickRandomP(arr []int) int{
	return arr[getRandomN(0, len(arr)-1)]
}

func PrimeNumbers() []int {
	list := make([]int, 0)
	cant := getRandomN(10, 15)
	i := 100
	for len(list) < cant {
		if isPrime(i){
			list = append(list, i)
		}
		i++
	}
	return list
}

func main() {
	println(pickRandomP(PrimeNumbers()))
}
