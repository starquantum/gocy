package common

// 计算两个字符串的相似度
func SimilarText(first, second string) float64 {
	if first == second {
		return 1
	}
	l1, l2 := len(first), len(second)
	sim := similarText(first, second, l1, l2)
	percent := float64(sim*200) / float64(l1+l2)
	return percent
}

func similarText(str1, str2 string, len1, len2 int) int {
	var sum, max int
	pos1, pos2 := 0, 0
	// Find the longest segment of the same section in two strings
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
				if l+1 > max {
					max = l + 1
					pos1 = i
					pos2 = j
				}
			}
		}
	}
	if sum = max; sum > 0 {
		if pos1 > 0 && pos2 > 0 {
			sum += similarText(str1, str2, pos1, pos2)
		}
		if (pos1+max < len1) && (pos2+max < len2) {
			s1 := []byte(str1)
			s2 := []byte(str2)
			sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
		}
	}
	return sum
}