package utils

func FilledCount(c int, b [][]string) (int, bool) {
	var f int

	for _, r := range b {
		for _, s := range r {
			if s != " " {
				f++
			}
		}
	}

	if f <= c-2 || f >= c+2 {
		return f, false
	}

	return f, true
}

func GophersExist(b [][]string) bool {
	for _, r := range b {
		for _, s := range r {
			if s == "g" {
				return true
			}
		}
	}

	return false
}

func FillBoardWith(item string, b [][]string) [][]string {
	return [][]string{}
}
