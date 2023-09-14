package insert_delete_get_random_o1

import "math/rand"

type RandomizedSet struct {
	mp map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mp: make(map[int]bool),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.mp[val]; ok {
		return false
	}
	s.mp[val] = true
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.mp[val]; !ok {
		return false
	}
	delete(s.mp, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	l := len(s.mp)
	randNumber := rand.Intn(l)
	counter := 0

	for k := range s.mp {
		if counter == randNumber {
			return k
		}
		counter++
	}

	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
