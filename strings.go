package commonCLI

import (
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// StringToStringsArr transform string of comma seperated words to array of strings
func StringToStringsArr(str string) (strArr []string) {
	for _, s := range strings.Split(str, ",") {
		strArr = append(strArr, s)
	}
	return strArr
}

// StringToIntArr transform string of comma seperated words to array of ints
func StringToIntArr(str string) (intArr []int) {
	for _, s := range strings.Split(str, ",") {
		num, _ := strconv.Atoi(s)
		intArr = append(intArr, num)
	}

	sort.Ints(intArr)
	return intArr
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

// IsStringInSlice returns TRUE is slice contains string and false if not
func IsStringInSlice(a string, list []string) bool {
	// We need that to not filter for empty list
	if len(list) > 0 {
		for _, b := range list {
			if b == a {
				return true
			}
		}
		return false
	}
	return true
}

// RemoveIntDuplicates removes duplicated elements from int array
func RemoveIntDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// RemoveStringDuplicates removes duplicated elements from strings array
func RemoveStringDuplicates(str []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range str {
		if encountered[str[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[str[v]] = true
			// Append to result slice.
			if str[v] != "" {
				result = append(result, str[v])
			}
		}
	}
	// Return the new slice.
	return result
}

// DeleteSlicefromSlice deletes one slice from another and removes duplicate objects
func DeleteSlicefromSlice(slice, delete []int) []int {
	for _, d := range delete {
		for i := len(slice) - 1; i >= 0; i-- {
			if slice[i] == d {
				slice = append(slice[:i], slice[i+1:]...)
			}
		}
	}

	return RemoveIntDuplicates(slice)
}

//isRegexMatch checks if our string matches regex
func isRegexMatch(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)

	return match
}

//isStringInt checks if given string is an integer and also returns int as second value
func isStringInt(id string) (bool, int) {
	val, err := strconv.Atoi(id)

	if err != nil {
		return false, 0
	}

	return true, val
}

//isTypeExpected checks if our object is of expected type
func isTypeExpected(s, expected interface{}) bool {

	if reflect.TypeOf(s) == reflect.TypeOf(expected) {
		return true
	}

	return false
}
