package reverse

func String(word string) string {
	var res string

	for _, char := range word {
		res = string(char) + res
	}

	return res
}
