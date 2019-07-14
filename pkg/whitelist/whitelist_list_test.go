package whitelist

import (
	"testing"
)

func TestListInsert(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	list := NewList()

	ok := list.Insert(ip)
	if !ok {
		t.Errorf("TestInsert failed to insert IP %s:\n", ip)
	}

	if !list.Contain(ip) {
		t.Errorf("TestInsert failed to retrieve IP %s:\n", ip)
	}
}

func TestListDelete(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	list := NewList()

	list.Insert(ip)

	ok := list.Delete(ip)
	if !ok {
		t.Errorf("TestDelete failed to delete IP %s:\n", ip)
	}

	if list.Contain(ip) {
		t.Errorf("TestDelete failed to retrieve IP %s:\n", ip)
	}
}

func TestListContain(t *testing.T) {
	t.Parallel()

	ip := "127.0.0.1"
	list := NewList()

	if list.Contain(ip) {
		t.Errorf("TestContain retrieved not inserted IP %s:\n", ip)
	}

	list.Insert(ip)

	if !list.Contain(ip) {
		t.Errorf("TestContain failed to retrieve IP %s:\n", ip)
	}
}

func TestListLen(t *testing.T) {
	t.Parallel()

	ip1 := "127.0.0.1"
	ip2 := "127.0.0.2"
	list := NewList()

	list.Insert(ip1)

	if list.Len() != 1 {
		t.Errorf("TestLen returned wrong length (%d instead of 1)\n", list.Len())
	}

	list.Insert(ip1)
	if list.Len() != 2 {
		t.Errorf("TestLen returned wrong length after inserting duplicate (%d instead of 1)\n", list.Len())
	}

	list.Insert(ip2)
	if list.Len() != 3 {
		t.Errorf("TestLen returned wrong length (%d instead of 2)\n", list.Len())
	}
}
