package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)
func main() {
	args := os.Args
	n, _ := strconv.Atoi(args[1])

	countStart := time.Now()
	count := CountPrime(n)
	countEnd   := time.Now()

	fmt.Println("number of prime:", count)
	fmt.Println("execution time: ", countEnd.Sub(countStart).Milliseconds(), "ms")

	factorizationStart := time.Now()
	primeFactor := PrimeFactorization(n)
	factorizationEnd := time.Now()

	strPrimeFactorization := ""	//表示用に成形
	for key, value := range primeFactor {
		strPrimeFactorization += strconv.Itoa(key) + "^" + strconv.Itoa(value) + " * "
	}
	strPrimeFactorization = strPrimeFactorization[:len(strPrimeFactorization) - 3]

	fmt.Println("\nresult of prime factorization: ", strPrimeFactorization)
	fmt.Println("execution time: ", factorizationEnd.Sub(factorizationStart).Milliseconds(), "ms")

	//fmt.Println("prime numbers: ", PrimeList)
}