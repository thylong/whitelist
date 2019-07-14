package whitelist

import (
	"github.com/armon/go-radix"
)

// Tree implements a Radix Tree
// Based on the awesome library github.com/armon/go-radix
type Tree struct {
	content *radix.Tree
}

// Storage a datastructure usable for storage
type Storage interface {
	Insert(string) bool
	Delete(string) bool
	Contain(string) bool
	Len() int
}

// New create a new radix tree.
func New(kind string) (storage Storage) {
	switch kind {
	case "list":
		storage = &List{content: []string{}}
	case "hashmap":
		storage = &HashMap{content: map[string]bool{}}
	case "radix":
		storage = &Tree{content: radix.New()}
	default:
		storage = &Tree{content: radix.New()}
	}
	return storage
}

// Insert add an IP to the structure.
func (t *Tree) Insert(ip string) bool {
	_, ok := t.content.Insert(ip, 1)
	return !ok
}

// Delete delete an IP from the structure.
func (t *Tree) Delete(ip string) bool {
	_, ok := t.content.Delete(ip)
	return ok
}

// Contain returns true if an IP is in the structure.
func (t *Tree) Contain(ip string) bool {
	_, found := t.content.Get(ip)
	return found
}

// Len returns the number of IPs in the structure.
func (t *Tree) Len() int {
	length := t.content.Len()
	return length
}
