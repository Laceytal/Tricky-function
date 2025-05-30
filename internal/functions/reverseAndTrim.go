package functions

func reverseTrim(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	i := 0
	for i < len(runes) && runes[i] == '0' {
		i++
	}
	return string(runes[i:])
}
