package hw13hashtable

type chain struct {
	key  StringHashCode
	val  interface{}
	next *chain
}

type hashChain struct {
	// hash table size
	size int
	// hash table factor for rehash
	factor float64
	// threshold when need rehash
	threshold int
	// current size
	len int

	ar []*chain
}

func (h *hashChain) Put(key StringHashCode, val interface{}) error {
	hash := h.getHash(key)

	if h.ar[hash] == nil {
		h.ar[hash] = &chain{key: key, val: val}
		h.len++
		h.reHash()
		return nil
	}

	ch := h.ar[hash]
	for ch != nil {
		if ch.key == key {
			ch.val = val
			return nil
		}
		ch = ch.next
	}

	h.ar[hash] = &chain{key: key, val: val, next: h.ar[hash]}
	h.len++
	h.reHash()
	return nil
}

func (h *hashChain) Get(key StringHashCode) (interface{}, bool) {
	hash := h.getHash(key)

	ch := h.ar[hash]
	for ch != nil {
		if ch.key == key {
			return ch.val, true
		}
		ch = ch.next
	}
	return nil, false
}

func (h *hashChain) Remove(key StringHashCode) (interface{}, bool) {
	hash := h.getHash(key)

	var prev *chain
	ch := h.ar[hash]

	for ch != nil {
		if ch.key == key {
			if prev == nil {
				h.ar[hash] = ch.next
			} else {
				prev.next = ch.next
			}
			h.len--
			return ch.val, true
		}
		prev = ch
		ch = ch.next
	}

	return nil, false
}

func (h *hashChain) getHash(key StringHashCode) int {
	return key.GetHashCode() % h.size
}

func (h *hashChain) reHash() {
	if h.len <= h.threshold {
		return
	}

	h.size = h.size*2 + 1
	h.threshold = int(float64(h.size) * h.factor)
	ar := make([]*chain, h.size)

	for _, ch := range h.ar {
		c := ch
		for c != nil {
			hash := h.getHash(c.key)
			if ar[hash] == nil {
				ar[hash] = &chain{
					key: c.key,
					val: c.val,
				}
			} else {
				ar[hash] = &chain{
					key:  c.key,
					val:  c.val,
					next: ar[hash],
				}
			}
			c = c.next
		}
	}

	h.ar = ar
}

func NewHashChain() *hashChain {
	return NewHashChainWithSize(11)
}

func NewHashChainWithSize(size int) *hashChain {
	factor := 0.75
	return &hashChain{
		size:      size,
		factor:    factor,
		threshold: int(float64(size) * factor),
		ar:        make([]*chain, size),
	}
}
