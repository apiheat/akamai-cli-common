package commonCLI

import "strings"

// StringToStringsArr transform string of comma seperated word to array of words
func StringToStringsArr(str string) (strArr []string) {
	for _, s := range strings.Split(str, ",") {
		strArr = append(strArr, s)
	}
	return strArr
}

// StringsSlicesDifference give you diff between 2 strings slices
func StringsSlicesDifference(slice1 []string, slice2 []string) []string {
	var diff []string
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}
