// +build benchmark

package whitelist

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkListInsert(b *testing.B) {
	ip := "127.0.0.1"
	list := NewList()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list.Insert(ip)
	}
}

func BenchmarkHMInsert(b *testing.B) {
	ip := "127.0.0.1"
	hashmap := NewHashMap()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		hashmap.Insert(ip)
	}
}

func BenchmarkRadixInsert(b *testing.B) {
	ip := "127.0.0.1"
	trie := New("radix")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		trie.Insert(ip)
	}
}

func BenchmarkListContain(b *testing.B) {
	list := NewList()

	for i := 0; i < 1000000; i++ {
		space1 := rand.Intn(255)
		space2 := rand.Intn(255)
		space3 := rand.Intn(255)
		space4 := rand.Intn(255)

		ip := fmt.Sprintf("%d%d%d%d", space1, space2, space3, space4)

		list.Insert(ip)
	}
	list.Insert("203.0.113.195")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		found := list.Contain("203.0.113.195")
		if !found {
			b.FailNow()
		}
	}
}

func BenchmarkHMContain(b *testing.B) {
	hashmap := NewHashMap()

	for i := 0; i < 1000000; i++ {
		space1 := rand.Intn(255)
		space2 := rand.Intn(255)
		space3 := rand.Intn(255)
		space4 := rand.Intn(255)

		ip := fmt.Sprintf("%d%d%d%d", space1, space2, space3, space4)

		hashmap.Insert(ip)
	}
	hashmap.Insert("203.0.113.195")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		found := hashmap.Contain("203.0.113.195")
		if !found {
			b.FailNow()
		}
	}
}

func BenchmarkRadixContain(b *testing.B) {
	tree := New("radix")

	for i := 0; i < 1000000; i++ {
		space1 := rand.Intn(255)
		space2 := rand.Intn(255)
		space3 := rand.Intn(255)
		space4 := rand.Intn(255)

		ip := fmt.Sprintf("%d%d%d%d", space1, space2, space3, space4)

		tree.Insert(ip)
	}
	tree.Insert("203.0.113.195")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		found := tree.Contain("203.0.113.195")
		if !found {
			b.FailNow()
		}
	}
}
