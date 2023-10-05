package combination_sum

import "sort"

/*
python code:
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        def backtrack(combination, remain, curr, results):
            if remain == 0:
                results.append(list(combination))
                return

            for next_curr in range(curr, len(candidates)):

                if next_curr > curr \
                        and candidates[next_curr] == candidates[next_curr - 1]:
                    continue

                pick = candidates[next_curr]
                if remain - pick < 0:
                    break

                combination.append(pick)
                backtrack(combination, remain - pick, next_curr, results)
                combination.pop()

        candidates.sort()

        combination, results = [], []
        backtrack(combination, target, 0, results)

        return results
*/

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	sort.Ints(candidates)
	backtrack(&result, []int{}, candidates, target, 0)

	return result
}

func backtrack(result *[][]int, combination []int, candidates []int, remain int, curr int) {
	if remain == 0 {
		*result = append(*result, append([]int{}, combination...))
		return
	}

	for nextCurr := curr; nextCurr < len(candidates); nextCurr++ {
		if nextCurr > curr && candidates[nextCurr] == candidates[nextCurr-1] {
			continue
		}

		pick := candidates[nextCurr]

		if remain-pick < 0 {
			break
		}

		combination = append(combination, pick)
		backtrack(result, combination, candidates, remain-pick, nextCurr)
		combination = (combination)[:len(combination)-1]
	}
}
