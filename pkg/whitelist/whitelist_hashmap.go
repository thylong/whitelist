package whitelist

// HashMap wraps the map primitive
type HashMap struct {
	content map[string]bool
}

// NewHashMap create a new radix tree.
func NewHashMap() Storage {
	return &HashMap{content: map[string]bool{}}
}

// Insert add an IP to the structure.
func (hm *HashMap) Insert(ip string) bool {
	hm.content[ip] = true
	return true
}

// Delete delete an IP from the structure.
func (hm *HashMap) Delete(ip string) bool {
	delete(hm.content, ip)
	return true
}

// Contain returns true if an IP is in the structure.
func (hm *HashMap) Contain(ip string) bool {
	_, found := hm.content[ip]
	return found
}

// Len returns the number of IPs in the structure.
func (hm *HashMap) Len() int {
	return len(hm.content)
}
