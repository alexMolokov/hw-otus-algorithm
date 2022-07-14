package hw10tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAvlTreeRandom(t *testing.T) {
	lenData := 1000
	lenDelete, lenSearch := 100, 100

	m := make(map[int]bool, lenData)
	tree := AvlTree{}
	rand.Seed(time.Now().UnixNano())

	i := 0
	for i < lenData {
		el := rand.Intn(1000000)
		_, ok := m[el]
		if !ok {
			m[el] = true
			tree.Insert(el)
			i++
		}
	}

	tm := time.Now()
	for key := range m {
		tree.search(key)
		lenSearch--
		if lenSearch <= 0 {
			break
		}
	}
	fmt.Printf("search time rand = %d mks \n", time.Since(tm).Microseconds())

	tm = time.Now()
	for key := range m {
		tree.Remove(key)
		lenDelete--
		if lenDelete <= 0 {
			break
		}
	}
	fmt.Printf("delete time rand = %d mks \n", time.Since(tm).Microseconds())
	require.True(t, true)
}

func TestAvlTreeSorted(t *testing.T) {
	lenData := 1000
	lenDelete, lenSearch := 100, 100

	m := make(map[int]bool, lenData)
	tree := AvlTree{}

	i := 0
	for i < lenData {
		tree.Insert(i)
		m[i] = true
		i++
	}

	// search
	tm := time.Now()
	for key := range m {
		tree.search(key)
		lenSearch--
		if lenSearch <= 0 {
			break
		}
	}
	fmt.Printf("search time sorted = %d mks \n", time.Since(tm).Microseconds())

	tm = time.Now()
	for key := range m {
		tree.Remove(key)
		lenDelete--
		if lenDelete <= 0 {
			break
		}
	}
	fmt.Printf("delete time sorted = %d mks \n", time.Since(tm).Microseconds())

	require.True(t, true)
}
