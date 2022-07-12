package iteration

const repeatCount = 5

func repeated(char string, rep int) (repeated string) {
	for i := 0; i < rep; i++ {
		repeated += char
	}
	return repeated
}
