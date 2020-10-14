package reverse

// Reverse -
func Reverse(str string) string {
	ch := []rune(str)
	for i, j := 0, len(ch)-1; i < j; i, j = i+1, j-1 {
		ch[i] = ch[i] ^ ch[j]
		ch[j] = ch[i] ^ ch[j]
		ch[i] = ch[i] ^ ch[j]
	}

	return string(ch)
}
