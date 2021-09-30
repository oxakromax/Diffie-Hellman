package utils

import (
	"math/big"
	"math/rand"
	"time"
)

func GetRandomN(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
func GetRandomN64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}
func isPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

func PickRandomGorP(arr []int64) int64 {
	return arr[GetRandomN(0, len(arr)-1)]
}

func PrimeNumbers() []int64 {
	list := make([]int64, 0)
	cant := GetRandomN(4, 8)
	i := int64(200)
	for len(list) < cant {
		if isPrime(i) {
			list = append(list, i)
		}
		i++
	}
	return list
}

func isPrimitiveRoot(g, p int64) bool {
	G := big.NewInt(g)
	P := big.NewInt(p)
	P1 := big.NewInt(p - 1)
	return new(big.Int).Exp(G, P1, P).Int64() == int64(1)
}

func PrimitiveRoots(p int64) []int64 {
	list := make([]int64, 0)
	cant := GetRandomN(10, 15)
	i := int64(1)
	for len(list) < cant {
		if isPrimitiveRoot(i, p) {
			list = append(list, i)
		}
		i++
	}
	return list
}
