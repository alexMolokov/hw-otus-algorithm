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
	lenDelete, lenSearch := 100, 100

	m := make(map[int]bool, lenData)
	tree := SimpleTree{}
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

	// Все добавленные ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be")
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
	fmt.Printf("search time rand = %d mks \n", time.Since(tm).Microseconds())

	i = 0
	deleted := make([]int, lenDelete)
	tm = time.Now()
	for key := range m {
		tree.Remove(key)
		delete(m, key)
		deleted[i] = key
		i++
		if i >= lenDelete {
			break
		}
	}
	fmt.Printf("delete time rand = %d mks \n", time.Since(tm).Microseconds())

	for _, key := range deleted {
		require.False(t, tree.Search(key), fmt.Sprintf("удаленные ключи не могут находится в дереве %d, len=%d", key, len(m)))
	}

	// Все оставшиеся ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be in after delete")
	}
}

func TestSimpleTreeSorted(t *testing.T) {
	lenData := 1000
	lenDelete, lenSearch := 100, 100

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

	i = 0
	deleted := make([]int, lenDelete)

	tm = time.Now()
	for key := range m {
		tree.Remove(key)
		delete(m, key)
		deleted[i] = key
		i++
		if i >= lenDelete {
			break
		}
	}
	fmt.Printf("delete time sorted = %d mks \n", time.Since(tm).Microseconds())

	for _, key := range deleted {
		require.False(t, tree.Search(key), fmt.Sprintf("удаленные ключи не могут находится в дереве %d, len=%d", key, len(m)))
	}

	// Все оставшиеся ключи находятся в дереве
	for key := range m {
		require.True(t, tree.Search(key), "all different elements must be in after delete")
	}
}
