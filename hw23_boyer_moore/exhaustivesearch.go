package hw23

const patternNotFound = -1

type ExhaustiveSearch struct{}

// Search поиск индекса в тексте с которого начинается шаблон или -1.
func (s *ExhaustiveSearch) Search(text, pattern string) int {
	textR := []rune(text)
	patternR := []rune(pattern)

	lenText := len(textR)
	lenPattern := len(patternR)

	if lenPattern > lenText {
		return patternNotFound
	}

	to := lenText - lenPattern

	for i := 0; i <= to; i++ {
		j := lenPattern - 1
		for j >= 0 {
			if patternR[j] != textR[i+j] {
				break
			}
			j--
		}
		if j == -1 {
			return i
		}
	}

	return patternNotFound
}
