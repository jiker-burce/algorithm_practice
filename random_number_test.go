package algorithm_practice

import (
	rand2 "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func rn1(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand2.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func rn2(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = []byte(strconv.Itoa(rand.Intn(10)))[0]
	}
	return string(b)
}

func TestRandomNum1(t *testing.T) {
	//fmt.Println(rn1(88))
	//fmt.Println(rn2(88))
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
}

func BenchmarkNum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rn1(100)
	}
}

func BenchmarkNum2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rn2(100)
	}
}
