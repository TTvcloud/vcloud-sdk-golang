package util

import "strings"

func Split(str, sep string) []string {

	str = strings.TrimSpace(str)
	sep = strings.TrimSpace(sep)

	if str == "" {
		return []string{}
	}

	if sep == "" {
		return []string{str}
	}

	return strings.Split(str, sep)
}

func SplitToMap(str, sep string) map[string]bool {

	str = strings.TrimSpace(str)
	sep = strings.TrimSpace(sep)

	result := make(map[string]bool)

	if str == "" {
		return make(map[string]bool)
	}

	if sep == "" {
		result[str] = true
		return result
	}

	splitStrs := strings.Split(str, sep)

	for i := range splitStrs {
		result[splitStrs[i]] = true
	}

	return result
}
