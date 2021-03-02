package main

import (
	"math"
	"os"
	"strconv"
)

var PrimeList []int	//わかっている素数を格納するスライス

/*素数リストをテキストファイルから読み込む関数*/
/*リストとか使ってると余計に遅くなったから却下*/
/*func LoadPrimeList () {
	filename := "prime_list.txt"
	fp, err := os.OpenFile(filename, os.O_RDONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fsc := bufio.NewScanner(fp)
	fsc.Split(bufio.ScanWords)
	for fsc.Scan() {
		tmp, _ := strconv.Atoi(fsc.Text())
		PrimeList = append(PrimeList, tmp)
	}
}*/

/*intのスライスにある値が含まれているか調べる関数*/
/*func contains (s []int, n int) bool {
	for _, v := range s {
		if n == v {
			return true
		}
	}
	return false
}

func IsPrime(n int) bool{
	if contains(PrimeList, n) {
		return true
	}
	for _, v := range PrimeList {

		if n != v && n % v == 0 {
			return false
		}
		if n < v {	//nより大きい数は見る必要がない(リストが昇順ソートされている前提)
			break
		}
	}
	PrimeList = append(PrimeList, n)
	return true
}*/

/*素数リストをテキストファイルに書き込む関数*/
func SavePrimeList () {
	filename := "prime_list.txt"
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var strPrimeList string
	for _, v := range PrimeList {
		strPrimeList += strconv.Itoa(v) + " "
	}
	_, err = fp.WriteString(strPrimeList)
	if err != nil {
		panic(err)
	}
}

/*ある値が素数か調べる関数*/
func IsPrime (n int) bool {
	if n == 2 {
		return true
	} else if n < 2 || n & 1 == 0 {
		return false
	}

	i := 3
	for ; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

/*ある値以下の素数の数を数える関数*/
func CountPrime (n int) (count int) {
	count = 0
	if n < 2 {
		return 0
	}
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			count++
			PrimeList = append(PrimeList, i)
		}
	}
	SavePrimeList()
	return
}

/*ある値を素因数分解する関数*/
func PrimeFactorization (n int) (res map[int]int) {
	/*if n > PrimeList[len(PrimeList) - 1] { // リストに足りない素数を追加
		CountPrime(n)
	}*/
	res = map[int]int{}	//素因数を格納するマップ
	for _, v := range PrimeList {
		if n == 1 {	//
			break
		}
		if n % v == 0 {
			for n % v == 0 {
				n /= v
				res[v]++
			}
		}
	}
	return
}

