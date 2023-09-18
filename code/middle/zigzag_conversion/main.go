package zigzag_conversion

func convert(s string, numRows int) string {
	// case numRows = 1
	if numRows == 1 {
		return s
	}

	// смотри в рисунок в README.md
	result := make([]rune, 0, len(s))

	mapGroup := make(map[int][]rune, numRows)
	reverse := false
	rows := -1
	// 0 1 2 3 4 3 2 1 0 1 2 3 4 3 2 1
	for _, v := range s {
		// Манипуляция с индексом занесения
		if !reverse {
			rows++
			mapGroup[rows] = append(mapGroup[rows], v)
		} else {
			rows--
			mapGroup[rows] = append(mapGroup[rows], v)
		}

		if (reverse && rows == 0) || rows == numRows-1 {
			if reverse {
				reverse = false
			} else {
				reverse = true
			}
		}
	}

	// Заполним result
	for i := 0; i < numRows; i++ {
		result = append(result, mapGroup[i]...)
	}

	return string(result)
}
