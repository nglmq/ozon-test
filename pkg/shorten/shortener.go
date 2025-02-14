package url

import "math/rand"

func NewRandomURL() string {
	alias := make([]rune, 10)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "_")

	for i := range alias {
		alias[i] = chars[rand.Intn(len(chars))]
	}

	return string(alias)
}
