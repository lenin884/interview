package gas_station

// Time complexity O(n ^ 2)
// Space complexity O(1)
//func canCompleteCircuit(gas []int, cost []int) int {
//	// find the first index where gas[i] >= cost[i]
//	n := len(gas)
//	for i := 0; i < n; i++ {
//		if gas[i] >= cost[i] {
//			if canCompleteCircuitFromIndex(gas, cost, i) {
//				return i
//			}
//		}
//	}
//
//	return -1
//}
//
//func canCompleteCircuitFromIndex(gas []int, cost []int, index int) bool {
//	n := len(gas)
//	gasLeft := 0
//	for i := 0; i < n; i++ {
//		gasLeft += gas[(index+i)%n] - cost[(index+i)%n]
//		if gasLeft < 0 {
//			return false
//		}
//	}
//
//	return true
//}

// Time complexity O(n)
// Space complexity O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	totalGas, tankGas, start := 0, 0, 0

	for i := 0; i < n; i++ {
		totalGas += gas[i] - cost[i]
		tankGas += gas[i] - cost[i]
		if tankGas < 0 {
			start = i + 1
			tankGas = 0
		}
	}

	if totalGas < 0 {
		return -1
	}

	return start
}
