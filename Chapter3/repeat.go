package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func RepeatModified(character string) string {
	repeated := strings.Repeat(character, repeatCount)

	return repeated
}
