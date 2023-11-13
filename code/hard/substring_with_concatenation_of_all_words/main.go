package substring_with_concatenation_of_all_words

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)
	if len(words) == 0 || len(s) == 0 {
		return result
	}

	// Запоминаем длину первого слова
	wordLen := len(words[0])
	// Вычисляем всего количество символов во всех словах
	totalLen := wordLen * len(words)

	// Создаем словарь, где ключ - слово, а значение - количество вхождений
	wordFreqMap := make(map[string]int)
	// Заполняем словарь
	for _, word := range words {
		wordFreqMap[word]++
	}

	// Проходим по всем символам строки начиная с первого символа, который может быть началом слова
	for i := 0; i < wordLen; i++ {
		// Создаем указатели на начало и конец текущего слова
		left, right := i, i
		currentWordMap := make(map[string]int)

		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			if wordFreqMap[word] == 0 {
				left = right
				currentWordMap = make(map[string]int)
			} else {
				currentWordMap[word]++
				// Если количество вхождений текущего слова больше, чем количество вхождений в словаре, то
				// удаляем слова из начала строки, пока не сравняем количество вхождений
				for currentWordMap[word] > wordFreqMap[word] {
					// Удаляем слово из начала строки
					leftWord := s[left : left+wordLen]
					// Сдвигаем указатель начала строки
					left += wordLen
					// Уменьшаем количество вхождений слова в текущей строке
					currentWordMap[leftWord]--
				}

				if right-left == totalLen {
					result = append(result, left)
				}
			}
		}
	}

	return result
}
