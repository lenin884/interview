package text_justification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	line := make([]string, 0)
	lineLength := 0
	i := 0

	for i < len(words) {
		if lineLength+len(words[i])+len(line) > maxWidth {
			extraSpaces := maxWidth - lineLength
			spaces := extraSpaces / max(1, len(line)-1)
			remainder := extraSpaces % max(1, len(line)-1)

			for j := 0; j < max(1, len(line)-1); j++ {
				for k := 0; k < spaces; k++ {
					line[j] += " "
				}

				if remainder > 0 {
					line[j] += " "
					remainder--
				}
			}

			result = append(result, strings.Join(line, ""))
			line, lineLength = make([]string, 0), 0
		}

		line = append(line, words[i])
		lineLength += len(words[i])
		i++
	}

	// Handle last line
	lastLine := strings.Join(line, " ")
	trailSpace := maxWidth - len(lastLine)
	for i := 0; i < trailSpace; i++ {
		lastLine += " "
	}
	result = append(result, lastLine)

	return result
}
