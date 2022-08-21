package hw27

type StateMachine struct {
	states  [][]int
	symbols map[byte]int
}

// Search поиск только в англоязычном тексте (чтобы 1 символ - 1 байт).
func (s *StateMachine) Search(text, pattern string) int {
	if len(s.states) == 0 {
		s.initStates(pattern)
	}
	lenPattern := len(pattern)

	bytes := []byte(text)
	currentState := 0
	for i, v := range bytes {
		if symbol, ok := s.symbols[v]; ok {
			currentState = s.states[currentState][symbol]
			if currentState == lenPattern {
				return i - lenPattern + 1
			}
			continue
		}
		currentState = 0
	}

	return -1
}

func (s *StateMachine) initStates(pattern string) {
	s.symbols = make(map[byte]int)
	bytes := []byte(pattern)
	i := 0
	for _, v := range bytes {
		if _, ok := s.symbols[v]; !ok {
			s.symbols[v] = i
			i++
		}
	}

	s.states = make([][]int, len(bytes))
	for i := 0; i < len(s.states); i++ {
		s.states[i] = make([]int, len(s.symbols))
		for letterByte, index := range s.symbols {
			word := pattern[:i] + string(letterByte)
			k := i + 1
			for pattern[:k] != word[len(word)-k:] {
				k--
			}
			s.states[i][index] = k
		}
	}
}
