package substring_with_concatenation_of_all_words

import "testing"

func TestSubstring(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		s := "barfoofoofoothefoobarman"
		words := []string{"foo", "bar"}
		expected := []int{0, 15}

		actual := findSubstring(s, words)
		if len(actual) != len(expected) {
			t.Errorf("findSubstring(%s, %v) should be %v", s, words, expected)
		}

		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				t.Errorf("findSubstring(%s, %v) should be %v", s, words, expected)
			}
		}
	})

	t.Run("case 2", func(t *testing.T) {
		s := "wordgoodgoodgoodbestword"
		words := []string{"word", "good", "best", "word"}
		expected := []int{}

		actual := findSubstring(s, words)
		if len(actual) != len(expected) {
			t.Errorf("findSubstring(%s, %v) should be %v", s, words, expected)
		}

		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				t.Errorf("findSubstring(%s, %v) should be %v", s, words, expected)
			}
		}
	})

	t.Run("case 3", func(t *testing.T) {
		s := "wordgoodgoodgoodbestword"
		words := []string{"word", "good", "best", "good"}
		expected := []int{8}

		actual := findSubstring(s, words)
		if len(actual) != len(expected) {
			t.Errorf("findSubstring(%s, %v) should be %v", s, words, expected)
		}

		for i := 0; i < len(actual); i++ {
			if actual[i] != expected[i] {
				t.Errorf("findSubstring(%s, %v) should be %v", s, words, expected)
			}
		}
	})
}
