package hw23

type BoyerMoore struct {
	lastOccurrence map[rune]int
	text           []rune
	pattern        []rune
	shiftSuffix    []int
}

// Search поиск индекса в тексте с которого начинается шаблон или -1.
func (b *BoyerMoore) Search(text, pattern string) int {
	b.text = []rune(text)
	b.pattern = []rune(pattern)
	b.shiftSuffix = make([]int, len(b.pattern))

	if len(b.pattern) > len(b.text) {
		return patternNotFound
	}

	b.computeLastOccurrence()
	b.computeShiftSuffix()

	i := 0
	for i+len(b.pattern) <= len(b.text) {
		found := true
		for j := len(b.pattern) - 1; j >= 0; j-- {
			if b.pattern[j] != b.text[i+j] {
				found = false
				shift := b.getShift(i+j, j)
				i += shift
				break
			}
		}
		if found {
			return i
		}
	}

	return patternNotFound
}

// getShift возвращает максимальное безопасное смещение из нескольких эвристик.
func (b *BoyerMoore) getShift(i, j int) int {
	shift := b.getBadCharacterHeuristic(i, j)
	if shift == len(b.pattern) {
		return shift
	}
	shiftSuffix := b.getGoodSuffixHeuristic(i, j)

	if shiftSuffix > shift {
		return shiftSuffix
	}

	return shift
}

func (b *BoyerMoore) getGoodSuffixHeuristic(_, j int) int {
	if j+1 >= len(b.shiftSuffix) {
		return 1
	}
	return b.shiftSuffix[j+1]
}

// getBadCharacterHeuristic возвращает безопасное количество символов на которое можно сдвинуть образец
// минимум сдвиг равен 1 (отрицательный сдвиг не делаем)
// j - индекс символа в шаблоне где произошло несовпадение с символом текста.
func (b *BoyerMoore) getBadCharacterHeuristic(i, j int) int {
	v, ok := b.lastOccurrence[b.text[i]]
	if !ok {
		return j + 1
	}
	if j > v {
		return j - v
	}

	return 1
}

func (b *BoyerMoore) computeLastOccurrence() {
	// кладем в map так как не хочется заводить массив на 65536 символов
	b.lastOccurrence = make(map[rune]int, len(b.pattern))

	for i := 0; i < len(b.pattern)-1; i++ {
		b.lastOccurrence[b.pattern[i]] = i
	}
}

func (b *BoyerMoore) computeShiftSuffix() {
	lenPattern := len(b.pattern)
	for i := 0; i < lenPattern; i++ {
		b.shiftSuffix[i] = lenPattern
	}

	// ищем префиксы равные суффиксу и записываем сдвиг
	// baddba

	for i := 0; i <= lenPattern-2; i++ {
		found := true
		for k := i; k >= 0; k-- {
			if b.pattern[i-k] != b.pattern[lenPattern-1-k] {
				found = false
				break
			}
		}

		if found {
			b.shiftSuffix[i] = lenPattern - 1 - i
			continue
		}

		if i > 0 {
			b.shiftSuffix[i] = b.shiftSuffix[i-1]
		}
	}

	last := b.pattern[len(b.pattern)-1]
	currentPosition := len(b.pattern) - 2

	for currentPosition >= 0 {
		if b.pattern[currentPosition] != last {
			currentPosition += -1
			continue
		}

		shift := lenPattern - 1 - currentPosition
		if b.shiftSuffix[0] > shift {
			b.shiftSuffix[0] = shift
		}

		middle := len(b.shiftSuffix) / 2
		for i := 1; i < middle; i++ {
			if currentPosition-i < 0 {
				break
			}

			if b.pattern[currentPosition-i] != b.pattern[lenPattern-1-i] {
				break
			}

			if b.shiftSuffix[i] < shift {
				continue
			}

			b.shiftSuffix[i] = shift
		}

		currentPosition += -1
	}
}
