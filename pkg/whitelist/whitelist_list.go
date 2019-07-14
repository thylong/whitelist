package whitelist

// List wraps the map primitive
type List struct {
	content []string
}

// NewList create a new radix tree.
func NewList() Storage {
	return &List{content: []string{}}
}

// Insert add an IP to the structure.
func (l *List) Insert(ip string) bool {
	l.content = append(l.content, ip)
	return true
}

// Delete delete an IP from the structure.
func (l *List) Delete(ip string) bool {
	for i, v := range l.content {
		if v == ip {
			l.content = append(l.content[:i], l.content[i+1:]...)
		}
	}
	return true
}

// Contain returns true if an IP is in the structure.
func (l *List) Contain(ip string) bool {
	for _, v := range l.content {
		if v == ip {
			return true
		}
	}
	return false
}

// Len returns the number of IPs in the structure.
func (l *List) Len() int {
	return len(l.content)
}
