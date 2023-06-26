package hw13hashtable

type HashCoder interface {
	GetHashCode() int
}

type StringHashCode string

func (s *StringHashCode) GetHashCode() int {
	var hash int
	for _, ch := range *s {
		hash = 31*hash + int(ch)
	}
	return hash
}
