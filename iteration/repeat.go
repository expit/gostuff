package iteration

//Repeat repeats repeated string
func Repeat(character string, repetitions int) (repeated string) {
	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return
}
