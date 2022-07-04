package hw10tree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSimpleTreeRandom(t *testing.T) {
	lenData := 1000
	lenDelete := 100

	m := make(map[int]bool, lenData)
	tree := SimpleTree{}
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

	// Все добавленные ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be")
	}

	i = 0
	deleted := make([]int, lenDelete)
	for key := range m {
		tree.Remove(key)
		delete(m, key)
		deleted[i] = key
		i++
		if i >= lenDelete {
			break
		}
	}

	for _, key := range deleted {
		require.False(t, tree.Search(key), fmt.Sprintf("удаленные ключи не могут находится в дереве %d, len=%d", key, len(m)))
	}

	// Все оставшиеся ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be in after delete")
	}
	fmt.Printf("time = %d ms", time.Since(tm).Milliseconds())
}

func TestSimpleTreeSorted(t *testing.T) {
	tm := time.Now()
	lenData := 1000
	lenDelete := 100

	m := make(map[int]bool, lenData)
	tree := SimpleTree{}

	i := 0
	for i < lenData {
		tree.Insert(i)
		m[i] = true
		i++
	}

	// Все добавленные ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be")
	}

	i = 0
	deleted := make([]int, lenDelete)
	for key := range m {
		tree.Remove(key)
		delete(m, key)
		deleted[i] = key
		i++
		if i >= lenDelete {
			break
		}
	}

	for _, key := range deleted {
		require.False(t, tree.Search(key), fmt.Sprintf("удаленные ключи не могут находится в дереве %d, len=%d", key, len(m)))
	}

	// Все оставшиеся ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be in after delete")
	}

	fmt.Printf("time = %d ms", time.Since(tm).Milliseconds())
}
