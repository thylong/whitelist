package whitelist

import (
	"testing"
)

func TestHMInsert(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	hashmap := NewHashMap()

	ok := hashmap.Insert(ip)
	if !ok {
		t.Errorf("TestInsert failed to insert IP %s:\n", ip)
	}

	if !hashmap.Contain(ip) {
		t.Errorf("TestInsert failed to retrieve IP %s:\n", ip)
	}
}

func TestHMDelete(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	hashmap := NewHashMap()

	hashmap.Insert(ip)

	ok := hashmap.Delete(ip)
	if !ok {
		t.Errorf("TestDelete failed to delete IP %s:\n", ip)
	}

	if hashmap.Contain(ip) {
		t.Errorf("TestDelete failed to retrieve IP %s:\n", ip)
	}
}

func TestHMContain(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	hashmap := NewHashMap()

	if hashmap.Contain(ip) {
		t.Errorf("TestContain retrieved not inserted IP %s:\n", ip)
	}

	hashmap.Insert(ip)

	if !hashmap.Contain(ip) {
		t.Errorf("TestContain failed to retrieve IP %s:\n", ip)
	}
}

func TestHMLen(t *testing.T) {
	t.Parallel()

	ip1 := "127.0.0.1"
	ip2 := "127.0.0.2"
	hashmap := NewHashMap()

	hashmap.Insert(ip1)

	if hashmap.Len() != 1 {
		t.Errorf("TestLen returned wrong length (%d instead of 1)\n", hashmap.Len())
	}

	hashmap.Insert(ip1)
	if hashmap.Len() != 1 {
		t.Errorf("TestLen returned wrong length after inserting duplicate (%d instead of 1)\n", hashmap.Len())
	}

	hashmap.Insert(ip2)
	if hashmap.Len() != 2 {
		t.Errorf("TestLen returned wrong length (%d instead of 2)\n", hashmap.Len())
	}
}
