package trygofuzz

func Checker(target string) bool {
	if target == "panic" {
		panic("panic")
	}
	if len(target) == 0 {
		return false
	}
	c := target[len(target) - 1]
	if 'a' <= c && c <= 'z' {
		panic("Tha last charactor is a")
	}
	return target == "message"
}
