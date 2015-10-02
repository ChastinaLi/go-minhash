package minhash

import (
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"testing"
	"reflect"
)

var h1, h2 Hash64
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	hash64a := fnv.New64a()
	h1 = func(b []byte) uint64 {
		hash64a.Reset()
		hash64a.Write(b)
		return hash64a.Sum64()
	}
  hash64 := fnv.New64()
  h2 = func(b []byte) uint64 {
    hash64.Reset()
    hash64.Write(b)
    return hash64.Sum64()
  }
}

func randSeq(n int) string {
  b := make([]rune, n)
  for i := range b {
      b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}

func data(size int) {
	for i := range size {
		Generate_hash(randSeq(10))
	}
}

func hashing(mh *MinWise, start, end int, data [][]byte) {
	fmt.Println(data)
	for i := start; i < end; i++ {
		mh.Push(data[i])
	}
}

func benchmark(minhashSize int, b *testing.B) {
	if b.N < 10 {
		fmt.Printf("\n")
		return
	}
	// Data is a set of unique values
	d := data(b.N)
	// a and b are two subsets of data with some overlaps
	a_start, a_end := 0, int(float64(b.N)*0.65)
	b_start, b_end := int(float64(b.N)*0.35), b.N

	m1 := NewMinWise(h1, h2, minhashSize)
	m2 := NewMinWise(h1, h2, minhashSize)

	b.ResetTimer()
	hashing(m1, a_start, a_end, d)
	hashing(m2, b_start, b_end, d)

	est := m1.Similarity(m2)
	act := float64(a_end-b_start) / float64(b_end-a_start)
	err := math.Abs(act - est)
	fmt.Printf("Data size: %8d, ", b.N)
	fmt.Printf("Real: %.8f, ", act)
	fmt.Printf("Estimated: %.8f, ", est)
	fmt.Printf("Error: %.8f\n", err)

	bytearray, _ := m1.Serialize()
	m1_deserialized, _ := Deserialize(bytearray)
	fmt.Println(reflect.DeepEqual(m1_deserialized.minimums , m1.minimums))
}

func BenchmarkMinWise64(b *testing.B) {
	benchmark(64, b)
}
//
// func BenchmarkMinWise128(b *testing.B) {
// 	benchmark(128, b)
// }
//
// func BenchmarkMinWise256(b *testing.B) {
// 	benchmark(256, b)
// }
//
// func BenchmarkMinWise512(b *testing.B) {
// 	benchmark(512, b)
// }
