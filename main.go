package main

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

func getRandomN(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
func isPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

func pickRandomGorP(arr []int) int {
	return arr[getRandomN(0, len(arr)-1)]
}

func primeNumbers() []int {
	list := make([]int, 0)
	cant := getRandomN(2, 6)
	i := 5
	for len(list) < cant {
		if isPrime(i) {
			list = append(list, i)
		}
		i++
	}
	return list
}

func isPrimitiveRoot(g, p int) bool {
	return (g^(p-1))%p == 1
}

func primitiveRoots(p int) []int {
	list := make([]int, 0)
	cant := getRandomN(1, 5)
	i := 1
	for len(list) < cant {
		if isPrimitiveRoot(i, p) {
			list = append(list, i)
		}
		i++
	}
	return list
}

type Diffie struct {
	privateKey, comunnicatorKey float64
	PublicKey, G, P             float64
}

func (d *Diffie) PrivateKey() float64 {
	return d.privateKey
}

func (d *Diffie) setComunnicatorKey(comunnicatorKey float64) {
	d.comunnicatorKey = comunnicatorKey
}

func (d *Diffie) config(p, g int) {
	d.P, d.G = float64(p), float64(g)
	d.privateKey = float64(getRandomN(1, p-1))
	d.PublicKey = math.Mod(math.Pow(d.G, d.privateKey), d.P) // A or B
}

func (d Diffie) getKey() float64 {
	return math.Mod(math.Pow(d.comunnicatorKey, d.PrivateKey()), d.P)
	//return (d.comunnicatorKey ^ d.privateKey) % d.P
}

//type server struct {
//	privateKey, clientKey int
//	PublicKey, G, P       int
//}
//
//func (s *server) configServer(p, g int) {
//	s.P, s.G = p, g
//	s.privateKey = getRandomN(1, p-1)
//	s.PublicKey = (s.G ^ s.privateKey) % s.P
//}

func main() {
	sucess := 0
	for i := 0; i < 10000; i++ {
		P := pickRandomGorP(primeNumbers())
		G := pickRandomGorP(primitiveRoots(P))
		client, server := new(Diffie), new(Diffie)
		client.config(P, G)
		server.config(P, G)
		client.setComunnicatorKey(server.PublicKey)
		server.setComunnicatorKey(client.PublicKey)
		if client.getKey() == server.getKey() {
			sucess++
		}
	}
	println(sucess)
	//text, _ := json.Marshal(client)
}
