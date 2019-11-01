package iteration

import "strings"

//Repeat takes a string and repeats it 5 times then returns it
func Repeat(value string, num int) string {
	return strings.Repeat(value, num)
}
