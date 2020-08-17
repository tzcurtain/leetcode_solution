package main

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int, len(s))
	for i := 0; i <= len(s)-10; i++ {
		nows := s[i : i+10]
		m[nows] = m[nows] + 1
	}
	var res []string
	for i := range m {
		if m[i] > 1 {
			res = append(res, i)
		}
	}
	return res
}
