package int_to_roman

func intToRoman(num int) string {
	//vn := [13]int{1, 4, 5, 9,
	//	10, 40, 50, 90,
	//	100, 400, 500, 900,
	//	1000}
	//vs := [13]string{"I", "IV", "V", "IX",
	//	"X", "XL", "L", "XC",
	//	"C", "CD", "D", "CM",
	//	"M"}
	//i := 12
	//r := ""
	//for true {
	//	tmp := num / vn[i]
	//	num %= vn[i]
	//	for tmp > 0 {
	//		r += vs[i]
	//		tmp--
	//	}
	//	i--
	//	if num < 1 {
	//		break
	//	}
	//}

	listInteger := []int{1, 4, 5, 9, 10, 40, 50,
		90, 100, 400, 500, 900, 1000}
	listRoman := []string{"I", "IV", "V", "IX", "X",
		"XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	index := 12
	result := ""
	for {
		tmp := num / listInteger[index]
		num %= listInteger[index]

		for tmp > 0 {
			result += listRoman[index]
			tmp--
		}

		index--

		if num < 1 {
			break
		}
	}

	return result
}
