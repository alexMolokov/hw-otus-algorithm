package hw23

import "unicode/utf8"

const patternNotFound = -1

type ExhaustiveSearch struct{}

// Search поиск индекса в тексте с которого начинается шаблон или -1.
func (s *ExhaustiveSearch) Search(text, pattern string) int {
	lenText := utf8.RuneCountInString(text)
	lenPattern := utf8.RuneCountInString(pattern)

	if lenPattern > lenText {
		return patternNotFound
	}

	textR := []rune(text)
	patternR := []rune(pattern)
	to := lenText - lenPattern

	for i := 0; i <= to; i++ {
		j := 0
		for j < lenPattern {
			if patternR[j] != textR[i+j] {
				break
			}
			j++
		}
		if j == lenPattern {
			return i
		}
	}

	return patternNotFound
}
