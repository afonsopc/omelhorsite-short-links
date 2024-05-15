package utils

import "math/rand"

func random_string(length int) string {
	chars := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var str []rune
	for i := 0; i < length; i++ {
		str = append(str, chars[rand.Intn(len(chars))])
	}
	return string(str)
}
