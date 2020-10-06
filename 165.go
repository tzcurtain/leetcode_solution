package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	l1, l2 := len(v1), len(v2)

	for i := 0; i < min(l1, l2); i++ {
		val1, _ := strconv.Atoi(v1[i])
		val2, _ := strconv.Atoi(v2[i])
		if val1 > val2 {
			return 1
		} else if val1 < val2 {
			return -1
		}
	}

	if l1 > l2 {
		for i := l2 + 1; i < l1; i++ {
			val, _ := strconv.Atoi(v1[i])
			if val != 0 {
				return 1
			}
		}
	} else if l1 < l2 {
		for i := l1 + 1; i < l2; i++ {
			val, _ := strconv.Atoi(v2[i])
			if val != 0 {
				return 1
			}
		}
		return -1
	}

	return 0
}
