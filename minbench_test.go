package minhash

import (
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"testing"
	"reflect"
	"time"
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

func data(size int) [][]byte {
	d := make([][]byte, size)
	for i := range d {
		d[i] = []byte(randSeq(5))
	}
	return d
}

func benchmark(minhashSize int, b *testing.B) {
	rand.Seed(time.Now().Unix())
	d := data(b.N)
	number := rand.Intn(len(d))

	b.ResetTimer()
	m1 := NewMinWise(h1, h2, minhashSize, d[0:number])
	m2 := NewMinWise(h1, h2, minhashSize, d)

	est := m1.Similarity(m2)
	act := float64(number+1) / float64(len(d))
	err := math.Abs(act - est)
	fmt.Printf("Error: %.8f\n", err)

	bytearray, _ := m1.Serialize()
	m1_deserialized, _ := Deserialize(bytearray)
	fmt.Println(reflect.DeepEqual(m1_deserialized.minimums , m1.minimums))
}

func BenchmarkMinWise64(b *testing.B) {
	benchmark(64, b)
}

func BenchmarkMinWise128(b *testing.B) {
	benchmark(128, b)
}

func BenchmarkMinWise256(b *testing.B) {
	benchmark(256, b)
}

func BenchmarkMinWise512(b *testing.B) {
	benchmark(512, b)
}
