package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

func ConcatSlice(sliceToConcat []byte) string {
	var concatenatedString string
	for _, v := range sliceToConcat {
		concatenatedString += string(v)
		concatenatedString += "-"
	}
	concatenatedString = concatenatedString[:len(concatenatedString)-1]
	return concatenatedString
}
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := range sliceToEncrypt {
		sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount)
	}
}
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
