package whitelist

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	tree := New()

	ok := tree.Insert(ip)
	if !ok {
		t.Errorf("TestInsert failed to insert IP %s:\n", ip)
	}

	if !tree.Contain(ip) {
		t.Errorf("TestInsert failed to retrieve IP %s:\n", ip)
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	tree := New()

	tree.Insert(ip)

	ok := tree.Delete(ip)
	if !ok {
		t.Errorf("TestDelete failed to delete IP %s:\n", ip)
	}

	if tree.Contain(ip) {
		t.Errorf("TestDelete failed to retrieve IP %s:\n", ip)
	}
}

func TestContain(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	tree := New()

	if tree.Contain(ip) {
		t.Errorf("TestContain retrieved not inserted IP %s:\n", ip)
	}

	tree.Insert(ip)

	if !tree.Contain(ip) {
		t.Errorf("TestContain failed to retrieve IP %s:\n", ip)
	}
}

func TestLen(t *testing.T) {
	t.Parallel()

	ip1 := "127.0.0.1"
	ip2 := "127.0.0.2"
	tree := New()

	tree.Insert(ip1)

	if tree.Len() != 1 {
		t.Errorf("TestLen returned wrong length (%d instead of 1)\n", tree.Len())
	}

	tree.Insert(ip1)
	if tree.Len() != 1 {
		t.Errorf("TestLen returned wrong length after inserting duplicate (%d instead of 1)\n", tree.Len())
	}

	tree.Insert(ip2)
	if tree.Len() != 2 {
		t.Errorf("TestLen returned wrong length (%d instead of 2)\n", tree.Len())
	}
}

func BenchmarkInsert(b *testing.B) {
	ip := "127.0.0.1"
	tree := New()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tree.Insert(ip)
	}
}
