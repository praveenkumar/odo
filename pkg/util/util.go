package util

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

// ConvertLabelsToSelector converts the given labels to selector
func ConvertLabelsToSelector(labels map[string]string) string {
	var selector string
	isFirst := true
	for k, v := range labels {
		if isFirst {
			isFirst = false
			if v == "" {
				selector = selector + fmt.Sprintf("%v", k)
			} else {
				selector = fmt.Sprintf("%v=%v", k, v)
			}
		} else {
			if v == "" {
				selector = selector + fmt.Sprintf(",%v", k)
			} else {
				selector = selector + fmt.Sprintf(",%v=%v", k, v)
			}
		}
	}
	return selector
}

// GenerateRandomString generates a random string of lower case characters of
// the given size
func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
