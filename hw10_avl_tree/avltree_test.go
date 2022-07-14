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
	lenDelete := 100

	m := make(map[int]bool, lenData)
	tree := AvlTree{}
	rand.Seed(time.Now().UnixNano())

	tm := time.Now()
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

	for key := range m {
		tree.Remove(key)
		lenDelete--
		if lenDelete <= 0 {
			break
		}
	}

	require.True(t, true)
	fmt.Printf("time = %d mks", time.Since(tm).Microseconds())
}

func TestAvlTreeSorted(t *testing.T) {
	tm := time.Now()
	lenData := 1000
	lenDelete := 100

	m := make(map[int]bool, lenData)
	tree := AvlTree{}

	i := 0
	for i < lenData {
		tree.Insert(i)
		m[i] = true
		i++
	}

	for key := range m {
		tree.Remove(key)
		lenDelete--
		if lenDelete <= 0 {
			break
		}
	}

	require.True(t, true)
	fmt.Printf("time = %d mks", time.Since(tm).Microseconds())
}
