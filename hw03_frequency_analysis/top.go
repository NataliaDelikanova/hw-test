package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	sl := strings.Fields(s)
	if len(sl) == 0 {
		return []string{}
	}

	m := map[string]int{}
	for _, el := range sl {
		m[el]++
	}

	mapValues := make([]int, 0, len(m))
	for _, val := range m {
		mapValues = append(mapValues, val)
	}
	sort.Slice(mapValues, func(i, j int) bool { return mapValues[i] > mapValues[j] })

	uniqValues := []int{}
	uniqValues = append(uniqValues, mapValues[0])
	for i := range mapValues {
		if i+1 < len(mapValues) && mapValues[i+1] != mapValues[i] {
			uniqValues = append(uniqValues, mapValues[i+1])
		}
	}

	resultTop := make([]string, 0, 10)
	first, last := 0, 0
	for v := range uniqValues {
		for i, el := range m {
			if len(resultTop) < 10 && el == uniqValues[v] {
				resultTop = append(resultTop, i)
				last++
			}
		}
		sort.Strings(resultTop[first:last])
		first = last
	}
	return resultTop
}
