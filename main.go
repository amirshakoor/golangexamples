package golangexamples

//importing Sir ehtesham's pakage
import (
	"github.com/ehteshamz/greetings"
)

//Concatenating slice
func ConcatSlice(sliceToConcat []byte) string {
	var concatenatedString string
	//iterating and concatenating slice
	for _, v := range sliceToConcat {
		concatenatedString += string(v)
		concatenatedString += "-"
	}
	//removing last - from string
	concatenatedString = concatenatedString[:len(concatenatedString)-1]
	return concatenatedString
}

//Encrypting slice
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	//iterating to encrypt slice by using ceaser cypher (value = value + 3)
	for i := range sliceToEncrypt {
		sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount)
	}
}

//Printing Sir Ehtesham's function
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
