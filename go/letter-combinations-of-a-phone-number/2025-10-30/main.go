func letterCombinations(digits string) []string {
	hashMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := make([]string, 0)

	for x := len(digits) - 1; x >= 0; x-- {
		digit := digits[x]
		stringMap := hashMap[digit]

		newResult := make([]string, 0, len(result)*3)

		for _, letter := range stringMap {
			if x == len(digits)-1 {
				newResult = append(newResult, string(letter))
			} else {
				for _, prev := range result {
					newResult = append(newResult, fmt.Sprintf("%s%s", string(letter), prev))
				}
			}
		}

		result = newResult
	}

	return result
}
