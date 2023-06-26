package hw13hashtable

type openAddressItem struct {
	key     HashCoder
	val     interface{}
	deleted bool
}

type hashOpenAddress struct {
	size      int
	factor    float64
	threshold int
	len       int
	ar        []*openAddressItem
	probing   func(hashCode, size, i int) int
}

func (h *hashOpenAddress) SetProbing(f func(hashCode, size, i int) int) {
	h.probing = f
}

func (h *hashOpenAddress) Put(key HashCoder, val interface{}) error {
	i := 0
	hash := h.getHash(key, h.size, i)

	ch := h.ar[hash]
	for {
		if ch == nil {
			h.ar[hash] = &openAddressItem{key: key, val: val}
			h.len++
			h.reHash()
			return nil
		}

		if ch.key == key {
			ch.val = val
			ch.deleted = false
			return nil
		}

		i++
		hash = h.getHash(key, h.size, i)
		ch = h.ar[hash]
	}
}

func (h *hashOpenAddress) Get(key HashCoder) (interface{}, bool) {
	i := 0
	hash := h.getHash(key, h.size, i)
	firstDeleted := -1

	ch := h.ar[hash]
	for ch != nil {
		if firstDeleted == -1 && ch.deleted {
			firstDeleted = hash
		}
		if ch.key == key && !ch.deleted {
			if firstDeleted > -1 {
				h.ar[firstDeleted] = &openAddressItem{key: ch.key, val: ch.val}
				h.ar[hash].deleted = true
			}

			return ch.val, true
		}

		i++
		hash = h.getHash(key, h.size, i)
		ch = h.ar[hash]
	}

	return nil, false
}

func (h *hashOpenAddress) Remove(key HashCoder) (interface{}, bool) {
	i := 0
	hash := h.getHash(key, h.size, i)

	ch := h.ar[hash]
	for ch != nil {
		if ch.key == key {
			h.ar[hash].deleted = true
			h.len--
			return ch.val, true
		}

		i++
		hash = h.getHash(key, h.size, i)
		ch = h.ar[hash]
	}

	return nil, false
}

func (h *hashOpenAddress) getHash(key HashCoder, size int, i int) int {
	return h.probing(key.GetHashCode(), size, i)
}

func (h *hashOpenAddress) reHash() {
	if h.len <= h.threshold {
		return
	}

	h.size = h.size * 2
	h.threshold = int(float64(h.size) * h.factor)

	ar := h.ar
	h.ar = make([]*openAddressItem, h.size)

	for _, ch := range ar {
		if ch != nil && !ch.deleted {
			_ = h.Put(ch.key, ch.val)
		}
	}
}

func NewHashOpenAddress() *hashOpenAddress {
	return NewHashOpenAddressWithSize(11)
}

func NewHashOpenAddressWithSize(size int) *hashOpenAddress {
	factor := 0.75
	return &hashOpenAddress{
		size:      size,
		factor:    factor,
		threshold: int(float64(size) * factor),
		probing: func(hashCode, size, i int) int {
			return (hashCode + 7*i + i*i) % size
		},
		ar: make([]*openAddressItem, size),
	}
}
