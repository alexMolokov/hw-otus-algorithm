package hw16

import "errors"

var ErrKeyNotFound = errors.New("key not found")

type node struct {
	child  [128]*node
	isLast bool
	value  interface{}
}

type bor struct {
	root *node
}

func (b *bor) Insert(key string, value interface{}) {
	current := b.root
	for _, code := range key {
		if current.child[code] == nil {
			current.child[code] = &node{}
		}

		current = current.child[code]
	}
	current.value = value
	current.isLast = true
}

func (b *bor) Search(key string) bool {
	_, err := b.Get(key)
	return err == nil
}

func (b *bor) StartWith(prefix string) bool {
	current := b.root
	for _, code := range prefix {
		if current.child[code] == nil {
			return false
		}
		current = current.child[code]
	}
	return true
}

func (b *bor) Get(key string) (interface{}, error) {
	current := b.root
	for _, code := range key {
		if current.child[code] == nil {
			return nil, ErrKeyNotFound
		}
		current = current.child[code]
	}

	if current.isLast {
		return current.value, nil
	}
	return nil, ErrKeyNotFound
}

func NewBor() *bor {
	return &bor{
		root: &node{},
	}
}
