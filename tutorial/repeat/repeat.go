package iteration

func Repeat (character string, count int) string {
	var repeated string;
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}