package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pref := strs[0]
	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			pref = ""
		} else {
			for j := 0; j < len(strs[i]); j++ {
				if j >= len(pref) || strs[i][j] != pref[j] {
					pref = pref[:j]
					break
				} else if j == len(strs[i])-1 {
					pref = strs[i]
					break
				}
			}
		}
	}
	return pref
}
