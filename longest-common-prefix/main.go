package main

func longestCommonPrefix(strs []string) string {
	for i := 1; i <= len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i <= len(strs[j]) {
				if strs[0][0:i] == strs[j][0:i] {
					continue
				} else {
					return strs[0][0 : i-1]
				}
			} else {
				return strs[0][0 : i-1]
			}

		}
	}

	if len(strs) == 1 {
		return strs[0]
	}

	return strs[0]

}
