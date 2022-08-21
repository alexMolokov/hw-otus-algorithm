package hw27

// PiBruteForce (рассмотрим только англоязычный алфавит в котором 1 буква - 1 байт).
func PiBruteForce(pattern string) []int {
	bytes := []byte(pattern)
	lenPattern := len(bytes)
	pi := make([]int, lenPattern+1)

	for i := 1; i <= lenPattern; i++ {
		for j := 1; j < i; j++ {
			if equal(bytes[0:j], bytes[i-j:i]) {
				pi[i] = j
			}
		}
	}

	return pi
}

func Pi(pattern string) []int {
	bytes := []byte(pattern)
	lenPattern := len(bytes)
	pi := make([]int, lenPattern+1)

	for i := 2; i <= lenPattern; i++ {
		index := pi[i-1]
		if bytes[index] == bytes[i-1] {
			pi[i] = pi[i-1] + 1
			continue
		}

		index = pi[index]
		for {
			if bytes[index] == bytes[i-1] {
				pi[i] = index + 1
				break
			}
			index = pi[index]
			if index == 0 {
				pi[i] = 0
				break
			}
		}
	}

	return pi
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Kmp рассматриваем только англоязычные буквы и считаем что  # отсутствует в алфавите.
func Kmp(text, pattern string) int {
	lenPattern := len(pattern)
	text = pattern + "#" + text
	pi := Pi(text)
	for i := 0; i < len(pi); i++ {
		if pi[i] == lenPattern {
			return i - (lenPattern + 1) - lenPattern
		}
	}
	return -1
}
