package trygofuzz

func Fuzz(data []byte) int {
	Checker(string(data))

	return 1
}
