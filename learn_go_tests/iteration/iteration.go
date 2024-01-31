package iteration

func Repeat(character string, freq int) string {
	var repeated string
	for i := 0; i < freq; i++ {
		repeated += character
	}
	return repeated
}
