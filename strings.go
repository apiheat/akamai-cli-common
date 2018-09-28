package commonCLI

import "strings"

// StringToStringsArr transform string of comma seperated word to array of words
func StringToStringsArr(str string) (strArr []string) {
	for _, s := range strings.Split(str, ",") {
		strArr = append(strArr, s)
	}
	return strArr
}
